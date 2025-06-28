# VkOHSurfaceCreateInfoOHOS(3) Manual Page

## Name

VkOHSurfaceCreateInfoOHOS - The parameters for surface creation on Open Harmony OS platform



## [](#_c_specification)C Specification

The `VkOHSurfaceCreateInfoOHOS` structure is defined as:

```c++
// Provided by VK_OHOS_surface
typedef struct VkOHSurfaceCreateInfoOHOS {
    VkStructureType             sType;
    const void*                 pNext;
    VkSurfaceCreateFlagsOHOS    flags;
    OHNativeWindow*             window;
} VkOHSurfaceCreateInfoOHOS;
```

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `window`: the pointer to a [OHNativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/OHNativeWindow.html) to associate the surface with.

or the equivalent

```c++
// Provided by VK_OHOS_surface
typedef VkOHSurfaceCreateInfoOHOS VkSurfaceCreateInfoOHOS;
```

## [](#_members)Members

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkOHSurfaceCreateInfoOHOS-sType-sType)VUID-VkOHSurfaceCreateInfoOHOS-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OH_SURFACE_CREATE_INFO_OHOS`
- [](#VUID-VkOHSurfaceCreateInfoOHOS-pNext-pNext)VUID-VkOHSurfaceCreateInfoOHOS-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkOHSurfaceCreateInfoOHOS-flags-zerobitmask)VUID-VkOHSurfaceCreateInfoOHOS-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_OHOS\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_OHOS_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceCreateFlagsOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCreateFlagsOHOS.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOHSurfaceCreateInfoOHOS)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0