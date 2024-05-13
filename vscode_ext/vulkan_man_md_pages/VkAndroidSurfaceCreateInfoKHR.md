# VkAndroidSurfaceCreateInfoKHR(3) Manual Page

## Name

VkAndroidSurfaceCreateInfoKHR - Structure specifying parameters of a
newly created Android surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAndroidSurfaceCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_android_surface
typedef struct VkAndroidSurfaceCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkAndroidSurfaceCreateFlagsKHR    flags;
    struct ANativeWindow*             window;
} VkAndroidSurfaceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `window` is a pointer to the [ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html) to
  associate the surface with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkAndroidSurfaceCreateInfoKHR-window-01248"
  id="VUID-VkAndroidSurfaceCreateInfoKHR-window-01248"></a>
  VUID-VkAndroidSurfaceCreateInfoKHR-window-01248  
  `window` **must** point to a valid Android
  [ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html)

Valid Usage (Implicit)

- <a href="#VUID-VkAndroidSurfaceCreateInfoKHR-sType-sType"
  id="VUID-VkAndroidSurfaceCreateInfoKHR-sType-sType"></a>
  VUID-VkAndroidSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR`

- <a href="#VUID-VkAndroidSurfaceCreateInfoKHR-pNext-pNext"
  id="VUID-VkAndroidSurfaceCreateInfoKHR-pNext-pNext"></a>
  VUID-VkAndroidSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkAndroidSurfaceCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkAndroidSurfaceCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkAndroidSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_android_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_android_surface.html),
[VkAndroidSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidSurfaceCreateFlagsKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAndroidSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAndroidSurfaceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
