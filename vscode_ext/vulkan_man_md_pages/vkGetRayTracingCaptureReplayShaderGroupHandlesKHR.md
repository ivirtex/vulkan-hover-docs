# vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(3) Manual Page

## Name

vkGetRayTracingCaptureReplayShaderGroupHandlesKHR - Query opaque capture
replay data for pipeline shader group handles



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the opaque capture data of shader groups in a ray tracing
pipeline, call:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
VkResult vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    firstGroup,
    uint32_t                                    groupCount,
    size_t                                      dataSize,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device containing the ray tracing pipeline.

- `pipeline` is the ray tracing pipeline object containing the shaders.

- `firstGroup` is the index of the first group to retrieve a handle for
  from the
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pGroups`
  array.

- `groupCount` is the number of shader handles to retrieve.

- `dataSize` is the size in bytes of the buffer pointed to by `pData`.

- `pData` is a pointer to an application-allocated buffer where the
  results will be written.

## <a href="#_description" class="anchor"></a>Description

On success, an array of `groupCount` shader handles will be written to
`pData`, with each element being of size
[VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleCaptureReplaySize`.

Once queried, this opaque data **can** be provided at pipeline creation
time (in a subsequent execution), using
[VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`pShaderGroupCaptureReplayHandle`,
as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-capture-replay"
target="_blank" rel="noopener">Ray Tracing Capture Replay</a>.

If `pipeline` was created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` and
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineLibraryGroupHandles"
target="_blank" rel="noopener">pipelineLibraryGroupHandles</a> feature
is enabled applications **can** query capture replay group handles from
that pipeline. The capture replay handle remains bitwise identical for
any `pipeline` which references the pipeline library. Group indices are
assigned as-if the pipeline was created without
`VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`.

Valid Usage

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-04620"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-04620"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-04620  
  `pipeline` **must** be a ray tracing pipeline

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-04051"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-04051"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-04051  
  `firstGroup` **must** be less than the number of shader groups in
  `pipeline`

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-03483"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-03483"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-03483  
  The sum of `firstGroup` and `groupCount` **must** be less than or
  equal to the number of shader groups in `pipeline`

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-03484"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-03484"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-03484  
  `dataSize` **must** be at least
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleCaptureReplaySize`
  × `groupCount`

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03606"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03606"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03606  
  `VkPhysicalDeviceRayTracingPipelineFeaturesKHR`::`rayTracingPipelineShaderGroupHandleCaptureReplay`
  **must** be enabled to call this function

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-03607"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-03607"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-03607  
  `pipeline` **must** have been created with a `flags` that included
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-07829"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-07829"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-07829  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineLibraryGroupHandles"
  target="_blank" rel="noopener">pipelineLibraryGroupHandles</a> feature
  is not enabled, `pipeline` **must** not have been created with
  `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-device-parameter"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-device-parameter"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parameter"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parameter"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pData-parameter"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pData-parameter"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-arraylength"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-arraylength"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

- <a
  href="#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parent"
  id="VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parent"></a>
  VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetRayTracingCaptureReplayShaderGroupHandlesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
