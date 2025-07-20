# vkGetPipelinePropertiesEXT(3) Manual Page

## Name

vkGetPipelinePropertiesEXT - Query pipeline properties



## [](#_c_specification)C Specification

To query the pipeline properties call:

```c++
// Provided by VK_EXT_pipeline_properties
VkResult vkGetPipelinePropertiesEXT(
    VkDevice                                    device,
    const VkPipelineInfoEXT*                    pPipelineInfo,
    VkBaseOutStructure*                         pPipelineProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the pipeline.
- `pPipelineInfo` is a pointer to a [VkPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineInfoEXT.html) structure which describes the pipeline being queried.
- `pPipelineProperties` is a pointer to a [VkBaseOutStructure](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBaseOutStructure.html) structure in which the pipeline properties will be written.

## [](#_description)Description

To query a pipeline’s `pipelineIdentifier` pass a [VkPipelinePropertiesIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelinePropertiesIdentifierEXT.html) structure in `pPipelineProperties`. Each pipeline is associated with a `pipelineIdentifier` and the identifier is implementation specific.

Valid Usage

- [](#VUID-vkGetPipelinePropertiesEXT-pipeline-06738)VUID-vkGetPipelinePropertiesEXT-pipeline-06738  
  The `pipeline` member of `pPipelineInfo` **must** have been created with `device`
- [](#VUID-vkGetPipelinePropertiesEXT-pPipelineProperties-06739)VUID-vkGetPipelinePropertiesEXT-pPipelineProperties-06739  
  `pPipelineProperties` **must** be a valid pointer to a [VkPipelinePropertiesIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelinePropertiesIdentifierEXT.html) structure
- [](#VUID-vkGetPipelinePropertiesEXT-None-06766)VUID-vkGetPipelinePropertiesEXT-None-06766  
  The [`pipelinePropertiesIdentifier`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelinePropertiesIdentifier) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetPipelinePropertiesEXT-device-parameter)VUID-vkGetPipelinePropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelinePropertiesEXT-pPipelineInfo-parameter)VUID-vkGetPipelinePropertiesEXT-pPipelineInfo-parameter  
  `pPipelineInfo` **must** be a valid pointer to a valid [VkPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineInfoEXT.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_properties.html), [VkBaseOutStructure](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBaseOutStructure.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelinePropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0