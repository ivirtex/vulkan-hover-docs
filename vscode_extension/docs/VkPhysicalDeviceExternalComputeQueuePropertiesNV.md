# VkPhysicalDeviceExternalComputeQueuePropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceExternalComputeQueuePropertiesNV - Structure specifying hardware specific information and limits for VK\_NV\_external\_compute\_queue functionality



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExternalComputeQueuePropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_external_compute_queue
typedef struct VkPhysicalDeviceExternalComputeQueuePropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           externalDataSize;
    uint32_t           maxExternalQueues;
} VkPhysicalDeviceExternalComputeQueuePropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `externalDataSize` is the minimum size of the memory allocation that applications **can** pass to [vkGetExternalComputeQueueDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExternalComputeQueueDataNV.html).
- `maxExternalQueues` is the maximum number of external queues that an application can create.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExternalComputeQueuePropertiesNV-sType-sType)VUID-VkPhysicalDeviceExternalComputeQueuePropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_COMPUTE_QUEUE_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExternalComputeQueuePropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0