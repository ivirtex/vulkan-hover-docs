# vkGetPipelineExecutableInternalRepresentationsKHR(3) Manual Page

## Name

vkGetPipelineExecutableInternalRepresentationsKHR - Get internal representations of the pipeline executable



## [](#_c_specification)C Specification

Each pipeline executable **may** have one or more text or binary internal representations associated with it which are generated as part of the compile process. These **may** include the final shader assembly, a binary form of the compiled shader, or the shader compiler’s internal representation at any number of intermediate compile steps. To query the internal representations associated with a pipeline executable, call:

```c++
// Provided by VK_KHR_pipeline_executable_properties
VkResult vkGetPipelineExecutableInternalRepresentationsKHR(
    VkDevice                                    device,
    const VkPipelineExecutableInfoKHR*          pExecutableInfo,
    uint32_t*                                   pInternalRepresentationCount,
    VkPipelineExecutableInternalRepresentationKHR* pInternalRepresentations);
```

## [](#_parameters)Parameters

- `device` is the device that created the pipeline.
- `pExecutableInfo` describes the pipeline executable being queried.
- `pInternalRepresentationCount` is a pointer to an integer related to the number of internal representations available or queried, as described below.
- `pInternalRepresentations` is either `NULL` or a pointer to an array of [VkPipelineExecutableInternalRepresentationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableInternalRepresentationKHR.html) structures.

## [](#_description)Description

If `pInternalRepresentations` is `NULL`, then the number of internal representations associated with the pipeline executable is returned in `pInternalRepresentationCount`. Otherwise, `pInternalRepresentationCount` **must** point to a variable set by the application to the number of elements in the `pInternalRepresentations` array, and on return the variable is overwritten with the number of structures actually written to `pInternalRepresentations`. If `pInternalRepresentationCount` is less than the number of internal representations associated with the pipeline executable, at most `pInternalRepresentationCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available representations were returned.

While the details of the internal representations remain implementation-dependent, the implementation **should** order the internal representations in the order in which they occur in the compiled pipeline with the final shader assembly (if any) last.

Valid Usage

- [](#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipelineExecutableInfo-03276)VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipelineExecutableInfo-03276  
  The [`pipelineExecutableInfo`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineExecutableInfo) feature **must** be enabled
- [](#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03277)VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03277  
  The `pipeline` member of `pExecutableInfo` **must** have been created with `device`
- [](#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03278)VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03278  
  The `pipeline` member of `pExecutableInfo` **must** have been created with `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-device-parameter)VUID-vkGetPipelineExecutableInternalRepresentationsKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pExecutableInfo-parameter)VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pExecutableInfo-parameter  
  `pExecutableInfo` **must** be a valid pointer to a valid [VkPipelineExecutableInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableInfoKHR.html) structure
- [](#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentationCount-parameter)VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentationCount-parameter  
  `pInternalRepresentationCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentations-parameter)VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentations-parameter  
  If the value referenced by `pInternalRepresentationCount` is not `0`, and `pInternalRepresentations` is not `NULL`, `pInternalRepresentations` **must** be a valid pointer to an array of `pInternalRepresentationCount` [VkPipelineExecutableInternalRepresentationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableInternalRepresentationKHR.html) structures

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

[VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineExecutableInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableInfoKHR.html), [VkPipelineExecutableInternalRepresentationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutableInternalRepresentationKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelineExecutableInternalRepresentationsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0