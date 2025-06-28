# vkGetExternalComputeQueueDataNV(3) Manual Page

## Name

vkGetExternalComputeQueueDataNV - Retrieves data necessary for compatible external API initialization



## [](#_c_specification)C Specification

To query the implementation-specific data that must be passed to compatible external APIs during their initialization process call:

```c++
// Provided by VK_NV_external_compute_queue
void vkGetExternalComputeQueueDataNV(
    VkExternalComputeQueueNV                    externalQueue,
    VkExternalComputeQueueDataParamsNV*         params,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `externalQueue` is the [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html) to query the data for.
- `params` is a pointer to a [VkExternalComputeQueueDataParamsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueDataParamsNV.html) structure specifying parameters required for retrieval of the implementation-specific data.
- `pData` is a pointer to application-allocated memory in which the requested data will be returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetExternalComputeQueueDataNV-pData-08134)VUID-vkGetExternalComputeQueueDataNV-pData-08134  
  `pData` **must** be at least the size specified by the externalDataSize field in the [VkPhysicalDeviceExternalComputeQueuePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalComputeQueuePropertiesNV.html) structure

Valid Usage (Implicit)

- [](#VUID-vkGetExternalComputeQueueDataNV-externalQueue-parameter)VUID-vkGetExternalComputeQueueDataNV-externalQueue-parameter  
  `externalQueue` **must** be a valid [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html) handle
- [](#VUID-vkGetExternalComputeQueueDataNV-params-parameter)VUID-vkGetExternalComputeQueueDataNV-params-parameter  
  `params` **must** be a valid pointer to a [VkExternalComputeQueueDataParamsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueDataParamsNV.html) structure
- [](#VUID-vkGetExternalComputeQueueDataNV-pData-parameter)VUID-vkGetExternalComputeQueueDataNV-pData-parameter  
  `pData` **must** be a pointer value

## [](#_see_also)See Also

[VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [VkExternalComputeQueueDataParamsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueDataParamsNV.html), [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetExternalComputeQueueDataNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0