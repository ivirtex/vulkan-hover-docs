# VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX(3) Manual Page

## Name

VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX - Structure describing multiview limits that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX` structure is defined as:

```c++
// Provided by VK_NVX_multiview_per_view_attributes
typedef struct VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           perViewPositionAllComponents;
} VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`perViewPositionAllComponents` is `VK_TRUE` if the implementation supports per-view position values that differ in components other than the X component.

## [](#_description)Description

If the `VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX-sType-sType)VUID-VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX`

## [](#_see_also)See Also

[VK\_NVX\_multiview\_per\_view\_attributes](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_multiview_per_view_attributes.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0