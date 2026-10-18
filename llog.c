/*
 * =====================================================================================
 *
 *       Filename:  llog.c
 *
 *    Description: Logger
 *
 *        Version:  1.0
 *        Created:  06/09/2026 09:30:59 PM
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
#include "llog.h"

uint8_t llog_current_level = 2;
