# vkGetRayTracingShaderGroupStackSizeKHR(3) Manual Page

## Name

vkGetRayTracingShaderGroupStackSizeKHR - Query ray tracing pipeline
shader group shader stack size



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the pipeline stack size of shaders in a shader group in the ray
tracing pipeline, call:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
VkDeviceSize vkGetRayTracingShaderGroupStackSizeKHR(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    group,
    VkShaderGroupShaderKHR                      groupShader);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device containing the ray tracing pipeline.

- `pipeline` is the ray tracing pipeline object containing the shaders
  groups.

- `group` is the index of the shader group to query.

- `groupShader` is the type of shader from the group to query.

## <a href="#_description" class="anchor"></a>Description

The return value is the ray tracing pipeline stack size in bytes for the
specified shader as called from the specified shader group.

Valid Usage

- <a href="#VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-04622"
  id="VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-04622"></a>
  VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-04622  
  `pipeline` **must** be a ray tracing pipeline

- <a href="#VUID-vkGetRayTracingShaderGroupStackSizeKHR-group-03608"
  id="VUID-vkGetRayTracingShaderGroupStackSizeKHR-group-03608"></a>
  VUID-vkGetRayTracingShaderGroupStackSizeKHR-group-03608  
  The value of `group` must be less than the number of shader groups in
  `pipeline`

- <a href="#VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-03609"
  id="VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-03609"></a>
  VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-03609  
  The shader identified by `groupShader` in `group` **must** not be
  `VK_SHADER_UNUSED_KHR`

Valid Usage (Implicit)

- <a href="#VUID-vkGetRayTracingShaderGroupStackSizeKHR-device-parameter"
  id="VUID-vkGetRayTracingShaderGroupStackSizeKHR-device-parameter"></a>
  VUID-vkGetRayTracingShaderGroupStackSizeKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parameter"
  id="VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parameter"></a>
  VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a
  href="#VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-parameter"
  id="VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-parameter"></a>
  VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-parameter  
  `groupShader` **must** be a valid
  [VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderGroupShaderKHR.html) value

- <a href="#VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parent"
  id="VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parent"></a>
  VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderGroupShaderKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetRayTracingShaderGroupStackSizeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
