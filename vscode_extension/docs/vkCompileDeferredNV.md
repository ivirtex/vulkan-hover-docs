# vkCompileDeferredNV(3) Manual Page

## Name

vkCompileDeferredNV - Deferred compilation of shaders



## [](#_c_specification)C Specification

To compile a deferred shader in a pipeline call:

```c++
// Provided by VK_NV_ray_tracing
VkResult vkCompileDeferredNV(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    shader);
```

## [](#_parameters)Parameters

- `device` is the logical device containing the ray tracing pipeline.
- `pipeline` is the ray tracing pipeline object containing the shaders.
- `shader` is the index of the shader to compile.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCompileDeferredNV-pipeline-04621)VUID-vkCompileDeferredNV-pipeline-04621  
  `pipeline` **must** be a ray tracing pipeline
- [](#VUID-vkCompileDeferredNV-pipeline-02237)VUID-vkCompileDeferredNV-pipeline-02237  
  `pipeline` **must** have been created with `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV`
- [](#VUID-vkCompileDeferredNV-shader-02238)VUID-vkCompileDeferredNV-shader-02238  
  `shader` **must** not have been called as a deferred compile before

Valid Usage (Implicit)

- [](#VUID-vkCompileDeferredNV-device-parameter)VUID-vkCompileDeferredNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCompileDeferredNV-pipeline-parameter)VUID-vkCompileDeferredNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkCompileDeferredNV-pipeline-parent)VUID-vkCompileDeferredNV-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCompileDeferredNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0