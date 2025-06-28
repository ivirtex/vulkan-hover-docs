# VkPhysicalDevicePointClippingProperties(3) Manual Page

## Name

VkPhysicalDevicePointClippingProperties - Structure describing the point clipping behavior supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDevicePointClippingProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDevicePointClippingProperties {
    VkStructureType            sType;
    void*                      pNext;
    VkPointClippingBehavior    pointClippingBehavior;
} VkPhysicalDevicePointClippingProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance2
typedef VkPhysicalDevicePointClippingProperties VkPhysicalDevicePointClippingPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`pointClippingBehavior` is a [VkPointClippingBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPointClippingBehavior.html) value specifying the point clipping behavior supported by the implementation.

If the `VkPhysicalDevicePointClippingProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePointClippingProperties-sType-sType)VUID-VkPhysicalDevicePointClippingProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkPointClippingBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPointClippingBehavior.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePointClippingProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0