# VkSurfaceFullScreenExclusiveWin32InfoEXT(3) Manual Page

## Name

VkSurfaceFullScreenExclusiveWin32InfoEXT - Structure specifying additional creation parameters specific to Win32 fullscreen exclusive mode



## [](#_c_specification)C Specification

The `VkSurfaceFullScreenExclusiveWin32InfoEXT` structure is defined as:

```c++
// Provided by VK_KHR_win32_surface with VK_EXT_full_screen_exclusive
typedef struct VkSurfaceFullScreenExclusiveWin32InfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    HMONITOR           hmonitor;
} VkSurfaceFullScreenExclusiveWin32InfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `hmonitor` is the Win32 `HMONITOR` handle identifying the display to create the surface with.

## [](#_description)Description

Note

If `hmonitor` is invalidated (e.g. the monitor is unplugged) during the lifetime of a swapchain created with this structure, operations on that swapchain will return `VK_ERROR_OUT_OF_DATE_KHR`.

Note

It is the responsibility of the application to change the display settings of the targeted Win32 display using the appropriate platform APIs. Such changes **may** alter the surface capabilities reported for the created surface.

Valid Usage

- [](#VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-hmonitor-02673)VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-hmonitor-02673  
  `hmonitor` **must** be a valid `HMONITOR`

Valid Usage (Implicit)

- [](#VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-sType-sType)VUID-VkSurfaceFullScreenExclusiveWin32InfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_full\_screen\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_full_screen_exclusive.html), [VK\_KHR\_win32\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_win32_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceFullScreenExclusiveWin32InfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0