# VkExternalComputeQueueDataParamsNV(3) Manual Page

## Name

VkExternalComputeQueueDataParamsNV - Structure specifying parameters for implementation-specific data retrieval from the external compute queue



## [](#_c_specification)C Specification

The `VkExternalComputeQueueDataParamsNV` structure is defined as:

```c++
// Provided by VK_NV_external_compute_queue
typedef struct VkExternalComputeQueueDataParamsNV {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           deviceIndex;
} VkExternalComputeQueueDataParamsNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `deviceIndex` is the index of the device within a device group that the data is being queried for. This is ignored if device groups are not utilized.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExternalComputeQueueDataParamsNV-sType-sType)VUID-VkExternalComputeQueueDataParamsNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_COMPUTE_QUEUE_DATA_PARAMS_NV`
- [](#VUID-VkExternalComputeQueueDataParamsNV-pNext-pNext)VUID-VkExternalComputeQueueDataParamsNV-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetExternalComputeQueueDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExternalComputeQueueDataNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalComputeQueueDataParamsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0