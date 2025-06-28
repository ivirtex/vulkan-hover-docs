# vkGetDataGraphPipelinePropertiesARM(3) Manual Page

## Name

vkGetDataGraphPipelinePropertiesARM - Query properties of a data graph pipeline



## [](#_c_specification)C Specification

To query properties of a data graph pipeline, call:

```c++
// Provided by VK_ARM_data_graph
VkResult vkGetDataGraphPipelinePropertiesARM(
    VkDevice                                    device,
    const VkDataGraphPipelineInfoARM*           pPipelineInfo,
    uint32_t                                    propertiesCount,
    VkDataGraphPipelinePropertyQueryResultARM*  pProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the data graph pipeline.
- `pPipelineInfo` is a [VkDataGraphPipelineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineInfoARM.html) that describes the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) being queried.
- `propertiesCount` is the length of the `pProperties` array.
- `pProperties` is a pointer to an array of [VkDataGraphPipelinePropertyQueryResultARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyQueryResultARM.html) structures.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetDataGraphPipelinePropertiesARM-dataGraphPipeline-09869)VUID-vkGetDataGraphPipelinePropertiesARM-dataGraphPipeline-09869  
  The `dataGraphPipeline` member of `pPipelineInfo` **must** have been returned by a call to [vkCreateDataGraphPipelinesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelinesARM.html)
- [](#VUID-vkGetDataGraphPipelinePropertiesARM-dataGraphPipeline-09802)VUID-vkGetDataGraphPipelinePropertiesARM-dataGraphPipeline-09802  
  The `dataGraphPipeline` member of `pPipelineInfo` **must** have been created with `device`
- [](#VUID-vkGetDataGraphPipelinePropertiesARM-pProperties-09889)VUID-vkGetDataGraphPipelinePropertiesARM-pProperties-09889  
  There **must** not be two or more structures in the `pProperties` array with the same [VkDataGraphPipelinePropertyQueryResultARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyQueryResultARM.html)::`property`

Valid Usage (Implicit)

- [](#VUID-vkGetDataGraphPipelinePropertiesARM-device-parameter)VUID-vkGetDataGraphPipelinePropertiesARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDataGraphPipelinePropertiesARM-pPipelineInfo-parameter)VUID-vkGetDataGraphPipelinePropertiesARM-pPipelineInfo-parameter  
  `pPipelineInfo` **must** be a valid pointer to a valid [VkDataGraphPipelineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineInfoARM.html) structure
- [](#VUID-vkGetDataGraphPipelinePropertiesARM-pProperties-parameter)VUID-vkGetDataGraphPipelinePropertiesARM-pProperties-parameter  
  `pProperties` **must** be a valid pointer to an array of `propertiesCount` [VkDataGraphPipelinePropertyQueryResultARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyQueryResultARM.html) structures
- [](#VUID-vkGetDataGraphPipelinePropertiesARM-propertiesCount-arraylength)VUID-vkGetDataGraphPipelinePropertiesARM-propertiesCount-arraylength  
  `propertiesCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineInfoARM.html), [VkDataGraphPipelinePropertyQueryResultARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyQueryResultARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDataGraphPipelinePropertiesARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0