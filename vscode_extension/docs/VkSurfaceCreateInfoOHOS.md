# VkSurfaceCreateInfoOHOS(3) Manual Page

## Name

VkSurfaceCreateInfoOHOS - The parameters for surface creation on Open Harmony OS platform



## [](#_c_specification)C Specification

The `VkSurfaceCreateInfoOHOS` structure is defined as:

```c++
// Provided by VK_OHOS_surface
typedef struct VkSurfaceCreateInfoOHOS {
    VkStructureType             sType;
    const void*                 pNext;
    VkSurfaceCreateFlagsOHOS    flags;
    OHNativeWindow*             window;
} VkSurfaceCreateInfoOHOS;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `window`: is a pointer to a [OHNativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/OHNativeWindow.html) to associate the surface with.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSurfaceCreateInfoOHOS-sType-sType)VUID-VkSurfaceCreateInfoOHOS-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_CREATE_INFO_OHOS`
- [](#VUID-VkSurfaceCreateInfoOHOS-pNext-pNext)VUID-VkSurfaceCreateInfoOHOS-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSurfaceCreateInfoOHOS-flags-zerobitmask)VUID-VkSurfaceCreateInfoOHOS-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_OHOS\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_OHOS_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceCreateFlagsOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCreateFlagsOHOS.html), [vkCreateSurfaceOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSurfaceOHOS.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceCreateInfoOHOS).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0