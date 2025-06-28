# vkGetPipelineKeyKHR(3) Manual Page

## Name

vkGetPipelineKeyKHR - Generate the pipeline key from pipeline creation info



## [](#_c_specification)C Specification

To generate the key for a particular pipeline creation info, call:

```c++
// Provided by VK_KHR_pipeline_binary
VkResult vkGetPipelineKeyKHR(
    VkDevice                                    device,
    const VkPipelineCreateInfoKHR*              pPipelineCreateInfo,
    VkPipelineBinaryKeyKHR*                     pPipelineKey);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the pipeline object.
- `pPipelineCreateInfo` is `NULL` or a pointer to a [VkPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateInfoKHR.html) structure.
- `pPipelineKey` is a pointer to a [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html) structure in which the resulting key is returned.

## [](#_description)Description

If `pPipelineCreateInfo` is `NULL`, then the implementation **must** return the global key that applies to all pipelines. If the key obtained in this way changes between saving and restoring data obtained from [vkGetPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineBinaryDataKHR.html) in a different [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), then the application **must** assume that the restored data is invalid and cannot be passed to [vkCreatePipelineBinariesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineBinariesKHR.html). Otherwise the application **can** assume the data is still valid.

If `pPipelineCreateInfo` is not `NULL`, the key obtained functions as a method to compare two pipeline creation info structures. Implementations **may** not compare parts of a pipeline creation info which would not contribute to the final binary output. If a shader module identifier is used instead of a shader module, the `pPipelineKey` generated **must** be equal to the key generated when using the shader module from which the identifier was queried. If the content of two `pPipelineKey` are equal, pipelines created with the two `pPipelineCreateInfo->pNext` create infos **must** produce the same [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html) contents.

The pipeline key is distinct from pipeline binary key. Pipeline binary keys **can** only be obtained after compilation. The pipeline key is intended to optionally allow associating pipeline create info with multiple pipeline binary keys.

Valid Usage

- [](#VUID-vkGetPipelineKeyKHR-pNext-09605)VUID-vkGetPipelineKeyKHR-pNext-09605  
  The `pNext` chain of `pPipelineCreateInfo` **must** not set [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` to a value greater than `0`

Valid Usage (Implicit)

- [](#VUID-vkGetPipelineKeyKHR-device-parameter)VUID-vkGetPipelineKeyKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelineKeyKHR-pPipelineCreateInfo-parameter)VUID-vkGetPipelineKeyKHR-pPipelineCreateInfo-parameter  
  If `pPipelineCreateInfo` is not `NULL`, `pPipelineCreateInfo` **must** be a valid pointer to a valid [VkPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateInfoKHR.html) structure
- [](#VUID-vkGetPipelineKeyKHR-pPipelineKey-parameter)VUID-vkGetPipelineKeyKHR-pPipelineKey-parameter  
  `pPipelineKey` **must** be a valid pointer to a [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html), [VkPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelineKeyKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0