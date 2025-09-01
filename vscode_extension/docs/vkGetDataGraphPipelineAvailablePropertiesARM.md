# vkGetDataGraphPipelineAvailablePropertiesARM(3) Manual Page

## Name

vkGetDataGraphPipelineAvailablePropertiesARM - Query available properties of a data graph pipeline



## [](#_c_specification)C Specification

To query the properties of a data graph pipeline that can be obtained, call:

```c++
// Provided by VK_ARM_data_graph
VkResult vkGetDataGraphPipelineAvailablePropertiesARM(
    VkDevice                                    device,
    const VkDataGraphPipelineInfoARM*           pPipelineInfo,
    uint32_t*                                   pPropertiesCount,
    VkDataGraphPipelinePropertyARM*             pProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the data graph pipeline.
- `pPipelineInfo` is a [VkDataGraphPipelineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineInfoARM.html) that describes the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) being queried.
- `pPropertiesCount` is a pointer to an integer related to the number of properties available or queried, as described below.
- `pProperties` is either `NULL` or a pointer to an array of [VkDataGraphPipelinePropertyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyARM.html) enums.

## [](#_description)Description

If `pProperties` is `NULL`, then the number of properties associated with the data graph pipeline is returned in `pPropertiesCount`. Otherwise, `pPropertiesCount` **must** point to a variable set by the user to the number of elements in the `pProperties` array, and on return the variable is overwritten with the number of enums actually written to `pProperties`. If `pPropertiesCount` is less than the number of properties associated with the data graph pipeline, at most `pPropertiesCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available properties were returned.

Valid Usage

- [](#VUID-vkGetDataGraphPipelineAvailablePropertiesARM-dataGraphPipeline-09888)VUID-vkGetDataGraphPipelineAvailablePropertiesARM-dataGraphPipeline-09888  
  The `dataGraphPipeline` member of `pPipelineInfo` **must** have been created with `device`

Valid Usage (Implicit)

- [](#VUID-vkGetDataGraphPipelineAvailablePropertiesARM-device-parameter)VUID-vkGetDataGraphPipelineAvailablePropertiesARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDataGraphPipelineAvailablePropertiesARM-pPipelineInfo-parameter)VUID-vkGetDataGraphPipelineAvailablePropertiesARM-pPipelineInfo-parameter  
  `pPipelineInfo` **must** be a valid pointer to a valid [VkDataGraphPipelineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineInfoARM.html) structure
- [](#VUID-vkGetDataGraphPipelineAvailablePropertiesARM-pPropertiesCount-parameter)VUID-vkGetDataGraphPipelineAvailablePropertiesARM-pPropertiesCount-parameter  
  `pPropertiesCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetDataGraphPipelineAvailablePropertiesARM-pProperties-parameter)VUID-vkGetDataGraphPipelineAvailablePropertiesARM-pProperties-parameter  
  If the value referenced by `pPropertiesCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertiesCount` [VkDataGraphPipelinePropertyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyARM.html) values

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineInfoARM.html), [VkDataGraphPipelinePropertyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDataGraphPipelineAvailablePropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0