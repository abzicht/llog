/*
 * =====================================================================================
 *
 *       Filename:  llog.h
 *
 *    Description: Logger
 *
 *        Version:  1.0
 *        Created:  06/09/2026 09:27:41 PM
 *
 *         Author:  abzicht (abzicht@gmail.com)
 *  Copyright (C):  2026 Abzicht (https://github.com/abzicht)
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 * =====================================================================================
 */

#include <stdint.h>
#include <stdio.h>

extern uint8_t llog_current_level;

#define llog_level_set(level) (llog_current_level = level)

#ifdef DEBUG
#define llog_fh stdout
#else
#define llog_fh stderr
#endif
#define llog_conditional(level, fmt, ...) do { if (level >= llog_current_level) { fprintf(llog_fh, (fmt), ##__VA_ARGS__); }} while(0)

#define llog_debug(fmt, ...) llog_conditional(1, fmt, ##__VA_ARGS__)
#define llog_info(fmt, ...)  llog_conditional(2, fmt, ##__VA_ARGS__)
#define llog_warn(fmt, ...)  llog_conditional(3, fmt, ##__VA_ARGS__)
#define llog_error(fmt, ...) llog_conditional(4, fmt, ##__VA_ARGS__)
