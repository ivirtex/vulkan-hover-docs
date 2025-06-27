# VkSurfaceFullScreenExclusiveWin32InfoEXT(3) Manual Page

## Name

VkSurfaceFullScreenExclusiveWin32InfoEXT - Structure specifying
additional creation parameters specific to Win32 fullscreen exclusive
mode



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfaceFullScreenExclusiveWin32InfoEXT` structure is defined as:

``` c
// Provided by VK_KHR_win32_surface with VK_EXT_full_screen_exclusive
typedef struct VkSurfaceFullScreenExclusiveWin32InfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    HMONITOR           hmonitor;
} VkSurfaceFullScreenExclusiveWin32InfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `hmonitor` is the Win32 `HMONITOR` handle identifying the display to
  create the surface with.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If <code>hmonitor</code> is invalidated (e.g. the monitor is
unplugged) during the lifetime of a swapchain created with this
structure, operations on that swapchain will return
<code>VK_ERROR_OUT_OF_DATE_KHR</code>.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is the responsibility of the application to change the display
settings of the targeted Win32 display using the appropriate platform
APIs. Such changes <strong>may</strong> alter the surface capabilities
reported for the created surface.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-hmonitor-02673"
  id="VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-hmonitor-02673"></a>
  VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-hmonitor-02673  
  `hmonitor` **must** be a valid `HMONITOR`

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-sType-sType"
  id="VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-sType-sType"></a>
  VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VK_KHR_win32_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_surface.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceFullScreenExclusiveWin32InfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
