# VkWin32SurfaceCreateInfoKHR(3) Manual Page

## Name

VkWin32SurfaceCreateInfoKHR - Structure specifying parameters of a newly created Win32 surface object



## [](#_c_specification)C Specification

The `VkWin32SurfaceCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_win32_surface
typedef struct VkWin32SurfaceCreateInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    VkWin32SurfaceCreateFlagsKHR    flags;
    HINSTANCE                       hinstance;
    HWND                            hwnd;
} VkWin32SurfaceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `hinstance` is the Win32 `HINSTANCE` for the window to associate the surface with.
- `hwnd` is the Win32 `HWND` for the window to associate the surface with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWin32SurfaceCreateInfoKHR-hinstance-01307)VUID-VkWin32SurfaceCreateInfoKHR-hinstance-01307  
  `hinstance` **must** be a valid Win32 `HINSTANCE`
- [](#VUID-VkWin32SurfaceCreateInfoKHR-hwnd-01308)VUID-VkWin32SurfaceCreateInfoKHR-hwnd-01308  
  `hwnd` **must** be a valid Win32 `HWND`

Valid Usage (Implicit)

- [](#VUID-VkWin32SurfaceCreateInfoKHR-sType-sType)VUID-VkWin32SurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR`
- [](#VUID-VkWin32SurfaceCreateInfoKHR-pNext-pNext)VUID-VkWin32SurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkWin32SurfaceCreateInfoKHR-flags-zerobitmask)VUID-VkWin32SurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_KHR\_win32\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_win32_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkWin32SurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32SurfaceCreateFlagsKHR.html), [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWin32SurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWin32SurfaceCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0