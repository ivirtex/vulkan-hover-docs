# VkAndroidSurfaceCreateInfoKHR(3) Manual Page

## Name

VkAndroidSurfaceCreateInfoKHR - Structure specifying parameters of a newly created Android surface object



## [](#_c_specification)C Specification

The `VkAndroidSurfaceCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_android_surface
typedef struct VkAndroidSurfaceCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkAndroidSurfaceCreateFlagsKHR    flags;
    struct ANativeWindow*             window;
} VkAndroidSurfaceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `window` is a pointer to the [ANativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/ANativeWindow.html) to associate the surface with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAndroidSurfaceCreateInfoKHR-window-01248)VUID-VkAndroidSurfaceCreateInfoKHR-window-01248  
  `window` **must** point to a valid Android [ANativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/ANativeWindow.html)

Valid Usage (Implicit)

- [](#VUID-VkAndroidSurfaceCreateInfoKHR-sType-sType)VUID-VkAndroidSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR`
- [](#VUID-VkAndroidSurfaceCreateInfoKHR-pNext-pNext)VUID-VkAndroidSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkAndroidSurfaceCreateInfoKHR-flags-zerobitmask)VUID-VkAndroidSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_KHR\_android\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_android_surface.html), [VkAndroidSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidSurfaceCreateFlagsKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateAndroidSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAndroidSurfaceCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0