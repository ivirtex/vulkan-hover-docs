# VkWin32SurfaceCreateInfoKHR(3) Manual Page

## Name

VkWin32SurfaceCreateInfoKHR - Structure specifying parameters of a newly
created Win32 surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkWin32SurfaceCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_win32_surface
typedef struct VkWin32SurfaceCreateInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    VkWin32SurfaceCreateFlagsKHR    flags;
    HINSTANCE                       hinstance;
    HWND                            hwnd;
} VkWin32SurfaceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `hinstance` is the Win32 `HINSTANCE` for the window to associate the
  surface with.

- `hwnd` is the Win32 `HWND` for the window to associate the surface
  with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkWin32SurfaceCreateInfoKHR-hinstance-01307"
  id="VUID-VkWin32SurfaceCreateInfoKHR-hinstance-01307"></a>
  VUID-VkWin32SurfaceCreateInfoKHR-hinstance-01307  
  `hinstance` **must** be a valid Win32 `HINSTANCE`

- <a href="#VUID-VkWin32SurfaceCreateInfoKHR-hwnd-01308"
  id="VUID-VkWin32SurfaceCreateInfoKHR-hwnd-01308"></a>
  VUID-VkWin32SurfaceCreateInfoKHR-hwnd-01308  
  `hwnd` **must** be a valid Win32 `HWND`

Valid Usage (Implicit)

- <a href="#VUID-VkWin32SurfaceCreateInfoKHR-sType-sType"
  id="VUID-VkWin32SurfaceCreateInfoKHR-sType-sType"></a>
  VUID-VkWin32SurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR`

- <a href="#VUID-VkWin32SurfaceCreateInfoKHR-pNext-pNext"
  id="VUID-VkWin32SurfaceCreateInfoKHR-pNext-pNext"></a>
  VUID-VkWin32SurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkWin32SurfaceCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkWin32SurfaceCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkWin32SurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_win32_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_surface.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkWin32SurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32SurfaceCreateFlagsKHR.html),
[vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWin32SurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkWin32SurfaceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
