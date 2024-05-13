# VkPipelineCreateFlagBits(3) Manual Page

## Name

VkPipelineCreateFlagBits - Bitmask controlling how a pipeline is created



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in

- [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`flags`

- [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html)::`flags`

- [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`flags`

- [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`flags`

specify how a pipeline is created, and are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkPipelineCreateFlagBits {
    VK_PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT = 0x00000001,
    VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT = 0x00000002,
    VK_PIPELINE_CREATE_DERIVATIVE_BIT = 0x00000004,
  // Provided by VK_VERSION_1_1
    VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT = 0x00000008,
  // Provided by VK_VERSION_1_1
    VK_PIPELINE_CREATE_DISPATCH_BASE_BIT = 0x00000010,
  // Provided by VK_VERSION_1_3
    VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT = 0x00000100,
  // Provided by VK_VERSION_1_3
    VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT = 0x00000200,
  // Provided by VK_KHR_dynamic_rendering with VK_KHR_fragment_shading_rate
    VK_PIPELINE_CREATE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR = 0x00200000,
  // Provided by VK_KHR_dynamic_rendering with VK_EXT_fragment_density_map
    VK_PIPELINE_CREATE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT = 0x00400000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR = 0x00004000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR = 0x00008000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR = 0x00010000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR = 0x00020000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR = 0x00001000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR = 0x00002000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR = 0x00080000,
  // Provided by VK_NV_ray_tracing
    VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV = 0x00000020,
  // Provided by VK_KHR_pipeline_executable_properties
    VK_PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR = 0x00000040,
  // Provided by VK_KHR_pipeline_executable_properties
    VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR = 0x00000080,
  // Provided by VK_NV_device_generated_commands
    VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV = 0x00040000,
  // Provided by VK_KHR_pipeline_library
    VK_PIPELINE_CREATE_LIBRARY_BIT_KHR = 0x00000800,
  // Provided by VK_EXT_descriptor_buffer
    VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT = 0x20000000,
  // Provided by VK_EXT_graphics_pipeline_library
    VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT = 0x00800000,
  // Provided by VK_EXT_graphics_pipeline_library
    VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT = 0x00000400,
  // Provided by VK_NV_ray_tracing_motion_blur
    VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV = 0x00100000,
  // Provided by VK_EXT_attachment_feedback_loop_layout
    VK_PIPELINE_CREATE_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT = 0x02000000,
  // Provided by VK_EXT_attachment_feedback_loop_layout
    VK_PIPELINE_CREATE_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT = 0x04000000,
  // Provided by VK_EXT_opacity_micromap
    VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT = 0x01000000,
#ifdef VK_ENABLE_BETA_EXTENSIONS
  // Provided by VK_NV_displacement_micromap
    VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV = 0x10000000,
#endif
  // Provided by VK_EXT_pipeline_protected_access
    VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT = 0x08000000,
  // Provided by VK_EXT_pipeline_protected_access
    VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT = 0x40000000,
  // Provided by VK_VERSION_1_1
    VK_PIPELINE_CREATE_DISPATCH_BASE = VK_PIPELINE_CREATE_DISPATCH_BASE_BIT,
  // Provided by VK_KHR_dynamic_rendering with VK_KHR_fragment_shading_rate
    VK_PIPELINE_RASTERIZATION_STATE_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR = VK_PIPELINE_CREATE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR,
  // Provided by VK_KHR_dynamic_rendering with VK_EXT_fragment_density_map
    VK_PIPELINE_RASTERIZATION_STATE_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT = VK_PIPELINE_CREATE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT,
  // Provided by VK_KHR_device_group
    VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR = VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT,
  // Provided by VK_KHR_device_group
    VK_PIPELINE_CREATE_DISPATCH_BASE_KHR = VK_PIPELINE_CREATE_DISPATCH_BASE,
  // Provided by VK_EXT_pipeline_creation_cache_control
    VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT = VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT,
  // Provided by VK_EXT_pipeline_creation_cache_control
    VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT_EXT = VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT,
} VkPipelineCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT` specifies that the
  created pipeline will not be optimized. Using this flag **may** reduce
  the time taken to create the pipeline.

- `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT` specifies that the pipeline
  to be created is allowed to be the parent of a pipeline that will be
  created in a subsequent pipeline creation call.

- `VK_PIPELINE_CREATE_DERIVATIVE_BIT` specifies that the pipeline to be
  created will be a child of a previously created parent pipeline.

- `VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT` specifies that
  any shader input variables decorated as `ViewIndex` will be assigned
  values as if they were decorated as `DeviceIndex`.

- `VK_PIPELINE_CREATE_DISPATCH_BASE` specifies that a compute pipeline
  **can** be used with [vkCmdDispatchBase](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchBase.html) with
  a non-zero base workgroup.

- `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV` specifies that a pipeline is
  created with all shaders in the deferred state. Before using the
  pipeline the application **must** call
  [vkCompileDeferredNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCompileDeferredNV.html) exactly once on each
  shader in the pipeline before using the pipeline.

- `VK_PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR` specifies that the
  shader compiler should capture statistics for the pipeline executables
  produced by the compile process which **can** later be retrieved by
  calling
  [vkGetPipelineExecutableStatisticsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableStatisticsKHR.html).
  Enabling this flag **must** not affect the final compiled pipeline but
  **may** disable pipeline caching or otherwise affect pipeline creation
  time.

- `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`
  specifies that the shader compiler should capture the internal
  representations of pipeline executables produced by the compile
  process which **can** later be retrieved by calling
  [vkGetPipelineExecutableInternalRepresentationsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableInternalRepresentationsKHR.html).
  Enabling this flag **must** not affect the final compiled pipeline but
  **may** disable pipeline caching or otherwise affect pipeline creation
  time. When capturing IR from pipelines created with pipeline
  libraries, there is no guarantee that IR from libraries **can** be
  retrieved from the linked pipeline. Applications **should** retrieve
  IR from each library, and any linked pipelines, separately.

- `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` specifies that the pipeline
  **cannot** be used directly, and instead defines a *pipeline library*
  that **can** be combined with other pipelines using the
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  structure. This is available in ray tracing and graphics pipelines.

- `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
  specifies that an any-hit shader will always be present when an
  any-hit shader would be executed. A NULL any-hit shader is an any-hit
  shader which is effectively `VK_SHADER_UNUSED_KHR`, such as from a
  shader group consisting entirely of zeros.

- `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`
  specifies that a closest hit shader will always be present when a
  closest hit shader would be executed. A NULL closest hit shader is a
  closest hit shader which is effectively `VK_SHADER_UNUSED_KHR`, such
  as from a shader group consisting entirely of zeros.

- `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`
  specifies that a miss shader will always be present when a miss shader
  would be executed. A NULL miss shader is a miss shader which is
  effectively `VK_SHADER_UNUSED_KHR`, such as from a shader group
  consisting entirely of zeros.

- `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`
  specifies that an intersection shader will always be present when an
  intersection shader would be executed. A NULL intersection shader is
  an intersection shader which is effectively `VK_SHADER_UNUSED_KHR`,
  such as from a shader group consisting entirely of zeros.

- `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` specifies that
  triangle primitives will be skipped during traversal using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
  target="_blank" rel="noopener">pipeline trace ray</a> instructions.

- `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR` specifies that
  AABB primitives will be skipped during traversal using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
  target="_blank" rel="noopener">pipeline trace ray</a> instructions.

- `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
  specifies that the shader group handles **can** be saved and reused on
  a subsequent run (e.g. for trace capture and replay).

- `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV` specifies that the
  pipeline can be used in combination with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands</a>.

- `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` specifies
  that pipeline creation will fail if a compile is required for creation
  of a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object;
  `VK_PIPELINE_COMPILE_REQUIRED` will be returned by pipeline creation,
  and the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) will be set to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html).

- When creating multiple pipelines,
  `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT` specifies that
  control will be returned to the application if any individual pipeline
  returns a result which is not `VK_SUCCESS` rather than continuing to
  create additional pipelines.

- `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV` specifies that
  the pipeline is allowed to use `OpTraceRayMotionNV`.

- `VK_PIPELINE_CREATE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
  specifies that the pipeline will be used with a fragment shading rate
  attachment and dynamic rendering.

- `VK_PIPELINE_CREATE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`
  specifies that the pipeline will be used with a fragment density map
  attachment and dynamic rendering.

- `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` specifies that
  pipeline libraries being linked into this library **should** have link
  time optimizations applied. If this bit is omitted, implementations
  **should** instead perform linking as rapidly as possible.

- `VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`
  specifies that pipeline libraries should retain any information
  necessary to later perform an optimal link with
  `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT`.

- `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` specifies that a
  pipeline will be used with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorbuffers"
  target="_blank" rel="noopener">descriptor buffers</a>, rather than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets"
  target="_blank" rel="noopener">descriptor sets</a>.

- `VK_PIPELINE_CREATE_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` specifies
  that the pipeline **may** be used with an attachment feedback loop
  including color attachments. It is ignored if
  `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT` is set in
  `pDynamicStates`.

- `VK_PIPELINE_CREATE_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  specifies that the pipeline **may** be used with an attachment
  feedback loop including depth-stencil attachments. It is ignored if
  `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT` is set in
  `pDynamicStates`.

- `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT` specifies
  that the ray tracing pipeline **can** be used with acceleration
  structures which reference an opacity micromap array.

- `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`
  specifies that the ray tracing pipeline **can** be used with
  acceleration structures which reference a displacement micromap array.

- `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT` specifies that the
  pipeline **must** not be bound to a protected command buffer.

- `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT` specifies that the
  pipeline **must** not be bound to an unprotected command buffer.

It is valid to set both `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT` and
`VK_PIPELINE_CREATE_DERIVATIVE_BIT`. This allows a pipeline to be both a
parent and possibly a child in a pipeline hierarchy. See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pipeline-derivatives"
target="_blank" rel="noopener">Pipeline Derivatives</a> for more
information.

When an implementation is looking up a pipeline in a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-cache"
target="_blank" rel="noopener">pipeline cache</a>, if that pipeline is
being created using linked libraries, implementations **should** always
return an equivalent pipeline created with
`VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` if available,
whether or not that bit was specified.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Using <code>VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT</code>
(or not) when linking pipeline libraries is intended as a performance
tradeoff between host and device. If the bit is omitted, linking should
be faster and produce a pipeline more rapidly, but performance of the
pipeline on the target device may be reduced. If the bit is included,
linking may be slower but should produce a pipeline with device
performance comparable to a monolithically created pipeline. Using both
options can allow latency-sensitive applications to generate a
suboptimal but usable pipeline quickly, and then perform an optimal link
in the background, substituting the result for the suboptimally linked
pipeline as soon as it is available.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
