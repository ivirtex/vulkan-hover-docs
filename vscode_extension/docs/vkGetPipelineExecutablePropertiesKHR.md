# vkGetPipelineExecutablePropertiesKHR(3) Manual Page

## Name

vkGetPipelineExecutablePropertiesKHR - Get the executables associated with a pipeline



## [](#_c_specification)C Specification

When a pipeline is created, its state and shaders are compiled into zero or more device-specific executables, which are used when executing commands against that pipeline. To query the properties of these pipeline executables, call:

```c++
// Provided by VK_KHR_pipeline_executable_properties
VkResult vkGetPipelineExecutablePropertiesKHR(
    VkDevice                                    device,
    const VkPipelineInfoKHR*                    pPipelineInfo,
    uint32_t*                                   pExecutableCount,
    VkPipelineExecutablePropertiesKHR*          pProperties);
```

## [](#_parameters)Parameters

- `device` is the device that created the pipeline.
- `pPipelineInfo` describes the pipeline being queried.
- `pExecutableCount` is a pointer to an integer related to the number of pipeline executables available or queried, as described below.
- `pProperties` is either `NULL` or a pointer to an array of [VkPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutablePropertiesKHR.html) structures.

## [](#_description)Description

If `pProperties` is `NULL`, then the number of pipeline executables associated with the pipeline is returned in `pExecutableCount`. Otherwise, `pExecutableCount` **must** point to a variable set by the application to the number of elements in the `pProperties` array, and on return the variable is overwritten with the number of structures actually written to `pProperties`. If `pExecutableCount` is less than the number of pipeline executables associated with the pipeline, at most `pExecutableCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available properties were returned.

Valid Usage

- [](#VUID-vkGetPipelineExecutablePropertiesKHR-pipelineExecutableInfo-03270)VUID-vkGetPipelineExecutablePropertiesKHR-pipelineExecutableInfo-03270  
  The [`pipelineExecutableInfo`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineExecutableInfo) feature **must** be enabled
- [](#VUID-vkGetPipelineExecutablePropertiesKHR-pipeline-03271)VUID-vkGetPipelineExecutablePropertiesKHR-pipeline-03271  
  The `pipeline` member of `pPipelineInfo` **must** have been created with `device`

Valid Usage (Implicit)

- [](#VUID-vkGetPipelineExecutablePropertiesKHR-device-parameter)VUID-vkGetPipelineExecutablePropertiesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelineExecutablePropertiesKHR-pPipelineInfo-parameter)VUID-vkGetPipelineExecutablePropertiesKHR-pPipelineInfo-parameter  
  `pPipelineInfo` **must** be a valid pointer to a valid [VkPipelineInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineInfoKHR.html) structure
- [](#VUID-vkGetPipelineExecutablePropertiesKHR-pExecutableCount-parameter)VUID-vkGetPipelineExecutablePropertiesKHR-pExecutableCount-parameter  
  `pExecutableCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPipelineExecutablePropertiesKHR-pProperties-parameter)VUID-vkGetPipelineExecutablePropertiesKHR-pProperties-parameter  
  If the value referenced by `pExecutableCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pExecutableCount` [VkPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutablePropertiesKHR.html) structures

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineExecutablePropertiesKHR.html), [VkPipelineInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelineExecutablePropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0