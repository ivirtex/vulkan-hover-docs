# VkExternalComputeQueueDeviceCreateInfoNV(3) Manual Page

## Name

VkExternalComputeQueueDeviceCreateInfoNV - Structure specifying information about external compute queues relevant to device creation



## [](#_c_specification)C Specification

The `VkExternalComputeQueueDeviceCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_external_compute_queue
typedef struct VkExternalComputeQueueDeviceCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           reservedExternalQueues;
} VkExternalComputeQueueDeviceCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `reservedExternalQueues` is the maximum number of external queues an application **can** create at once. This **must** be less than or equal to the `maxExternalQueues` value reported by [VkPhysicalDeviceExternalComputeQueuePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalComputeQueuePropertiesNV.html)

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExternalComputeQueueDeviceCreateInfoNV-sType-sType)VUID-VkExternalComputeQueueDeviceCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_COMPUTE_QUEUE_DEVICE_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalComputeQueueDeviceCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0