# VkQueueFamilyOwnershipTransferPropertiesKHR(3) Manual Page

## Name

VkQueueFamilyOwnershipTransferPropertiesKHR - Structure describing queue family ownership transfer properties



## [](#_c_specification)C Specification

The [VkQueueFamilyOwnershipTransferPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyOwnershipTransferPropertiesKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_maintenance9
typedef struct VkQueueFamilyOwnershipTransferPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           optimalImageTransferToQueueFamilies;
} VkQueueFamilyOwnershipTransferPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `optimalImageTransferToQueueFamilies` is a bitmask of queue family indices that indicates which queue families belonging to the same logical device support implicitly acquiring optimal image resources owned by this queue family, without the resources' contents becoming undefined.

## [](#_description)Description

If this structure is included in the `pNext` chain of the [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html) structure passed to [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html), then it is filled with the queue family ownership properties for the specified queue family.

Valid Usage (Implicit)

- [](#VUID-VkQueueFamilyOwnershipTransferPropertiesKHR-sType-sType)VUID-VkQueueFamilyOwnershipTransferPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_OWNERSHIP_TRANSFER_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_maintenance9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance9.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyOwnershipTransferPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0