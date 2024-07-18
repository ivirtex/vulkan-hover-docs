# VkPipelineCreateFlagBits2KHR(3) Manual Page

## Name

VkPipelineCreateFlagBits2KHR - Bitmask controlling how a pipeline is
created



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)::`flags`,
specifying how a pipeline is created, are:

``` c
// Provided by VK_KHR_maintenance5
// Flag bits for VkPipelineCreateFlagBits2KHR
typedef VkFlags64 VkPipelineCreateFlagBits2KHR;
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_DISABLE_OPTIMIZATION_BIT_KHR = 0x00000001ULL;
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_ALLOW_DERIVATIVES_BIT_KHR = 0x00000002ULL;
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_DERIVATIVE_BIT_KHR = 0x00000004ULL;
// Provided by VK_EXT_legacy_dithering with (VK_KHR_dynamic_rendering or VK_VERSION_1_3) and VK_KHR_maintenance5
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_ENABLE_LEGACY_DITHERING_BIT_EXT = 0x400000000ULL;
// Provided by VK_KHR_maintenance5 with VK_VERSION_1_1 or VK_KHR_device_group
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR = 0x00000008ULL;
// Provided by VK_KHR_maintenance5 with VK_VERSION_1_1 or VK_KHR_device_group
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_DISPATCH_BASE_BIT_KHR = 0x00000010ULL;
// Provided by VK_KHR_maintenance5 with VK_NV_ray_tracing
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_DEFER_COMPILE_BIT_NV = 0x00000020ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_pipeline_executable_properties
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_CAPTURE_STATISTICS_BIT_KHR = 0x00000040ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_pipeline_executable_properties
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR = 0x00000080ULL;
// Provided by VK_KHR_maintenance5 with VK_VERSION_1_3 or VK_EXT_pipeline_creation_cache_control
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_KHR = 0x00000100ULL;
// Provided by VK_KHR_maintenance5 with VK_VERSION_1_3 or VK_EXT_pipeline_creation_cache_control
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_EARLY_RETURN_ON_FAILURE_BIT_KHR = 0x00000200ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_graphics_pipeline_library
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_LINK_TIME_OPTIMIZATION_BIT_EXT = 0x00000400ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_graphics_pipeline_library
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT = 0x00800000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_pipeline_library
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_LIBRARY_BIT_KHR = 0x00000800ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR = 0x00001000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_AABBS_BIT_KHR = 0x00002000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR = 0x00004000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR = 0x00008000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR = 0x00010000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR = 0x00020000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR = 0x00080000ULL;
// Provided by VK_KHR_maintenance5 with VK_NV_device_generated_commands
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_NV = 0x00040000ULL;
// Provided by VK_KHR_maintenance5 with VK_NV_ray_tracing_motion_blur
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_ALLOW_MOTION_BIT_NV = 0x00100000ULL;
// Provided by VK_KHR_maintenance5 with (VK_KHR_dynamic_rendering or VK_VERSION_1_3) and VK_KHR_fragment_shading_rate
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR = 0x00200000ULL;
// Provided by VK_KHR_maintenance5 with (VK_KHR_dynamic_rendering or VK_VERSION_1_3) and VK_EXT_fragment_density_map
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT = 0x00400000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_opacity_micromap
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT = 0x01000000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_attachment_feedback_loop_layout
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT = 0x02000000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_attachment_feedback_loop_layout
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT = 0x04000000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_pipeline_protected_access
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_NO_PROTECTED_ACCESS_BIT_EXT = 0x08000000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_pipeline_protected_access
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_PROTECTED_ACCESS_ONLY_BIT_EXT = 0x40000000ULL;
// Provided by VK_KHR_maintenance5 with VK_NV_displacement_micromap
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV = 0x10000000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_descriptor_buffer
static const VkPipelineCreateFlagBits2KHR VK_PIPELINE_CREATE_2_DESCRIPTOR_BUFFER_BIT_EXT = 0x20000000ULL;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_CREATE_2_DISABLE_OPTIMIZATION_BIT_KHR` specifies that the
  created pipeline will not be optimized. Using this flag **may** reduce
  the time taken to create the pipeline.

- `VK_PIPELINE_CREATE_2_ALLOW_DERIVATIVES_BIT_KHR` specifies that the
  pipeline to be created is allowed to be the parent of a pipeline that
  will be created in a subsequent pipeline creation call.

- `VK_PIPELINE_CREATE_2_DERIVATIVE_BIT_KHR` specifies that the pipeline
  to be created will be a child of a previously created parent pipeline.

- `VK_PIPELINE_CREATE_2_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR` specifies
  that any shader input variables decorated as `ViewIndex` will be
  assigned values as if they were decorated as `DeviceIndex`.

- `VK_PIPELINE_CREATE_2_DISPATCH_BASE_BIT_KHR` specifies that a compute
  pipeline **can** be used with
  [vkCmdDispatchBase](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchBase.html) with a non-zero base
  workgroup.

- `VK_PIPELINE_CREATE_2_DEFER_COMPILE_BIT_NV` specifies that a pipeline
  is created with all shaders in the deferred state. Before using the
  pipeline the application **must** call
  [vkCompileDeferredNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCompileDeferredNV.html) exactly once on each
  shader in the pipeline before using the pipeline.

- `VK_PIPELINE_CREATE_2_CAPTURE_STATISTICS_BIT_KHR` specifies that the
  shader compiler should capture statistics for the pipeline executables
  produced by the compile process which **can** later be retrieved by
  calling
  [vkGetPipelineExecutableStatisticsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableStatisticsKHR.html).
  Enabling this flag **must** not affect the final compiled pipeline but
  **may** disable pipeline caching or otherwise affect pipeline creation
  time.

- `VK_PIPELINE_CREATE_2_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`
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

- `VK_PIPELINE_CREATE_2_LIBRARY_BIT_KHR` specifies that the pipeline
  **cannot** be used directly, and instead defines a *pipeline library*
  that **can** be combined with other pipelines using the
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  structure. This is available in ray tracing and graphics pipelines.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
  specifies that an any-hit shader will always be present when an
  any-hit shader would be executed. A NULL any-hit shader is an any-hit
  shader which is effectively `VK_SHADER_UNUSED_KHR`, such as from a
  shader group consisting entirely of zeros.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`
  specifies that a closest hit shader will always be present when a
  closest hit shader would be executed. A NULL closest hit shader is a
  closest hit shader which is effectively `VK_SHADER_UNUSED_KHR`, such
  as from a shader group consisting entirely of zeros.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`
  specifies that a miss shader will always be present when a miss shader
  would be executed. A NULL miss shader is a miss shader which is
  effectively `VK_SHADER_UNUSED_KHR`, such as from a shader group
  consisting entirely of zeros.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`
  specifies that an intersection shader will always be present when an
  intersection shader would be executed. A NULL intersection shader is
  an intersection shader which is effectively `VK_SHADER_UNUSED_KHR`,
  such as from a shader group consisting entirely of zeros.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` specifies
  that triangle primitives will be skipped during traversal using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
  target="_blank" rel="noopener">pipeline trace ray</a> instructions.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_AABBS_BIT_KHR` specifies that
  AABB primitives will be skipped during traversal using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
  target="_blank" rel="noopener">pipeline trace ray</a> instructions.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
  specifies that the shader group handles **can** be saved and reused on
  a subsequent run (e.g. for trace capture and replay).

- `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_NV` specifies that the
  pipeline can be used in combination with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands</a>.

- `VK_PIPELINE_CREATE_2_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_KHR`
  specifies that pipeline creation will fail if a compile is required
  for creation of a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object;
  `VK_PIPELINE_COMPILE_REQUIRED` will be returned by pipeline creation,
  and the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) will be set to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html).

- When creating multiple pipelines,
  `VK_PIPELINE_CREATE_2_EARLY_RETURN_ON_FAILURE_BIT_KHR` specifies that
  control will be returned to the application if any individual pipeline
  returns a result which is not `VK_SUCCESS` rather than continuing to
  create additional pipelines.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_ALLOW_MOTION_BIT_NV` specifies that
  the pipeline is allowed to use `OpTraceRayMotionNV`.

- `VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
  specifies that the pipeline will be used with a fragment shading rate
  attachment.

- `VK_PIPELINE_CREATE_2_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`
  specifies that the pipeline will be used with a fragment density map
  attachment.

- `VK_PIPELINE_CREATE_2_LINK_TIME_OPTIMIZATION_BIT_EXT` specifies that
  pipeline libraries being linked into this library **should** have link
  time optimizations applied. If this bit is omitted, implementations
  **should** instead perform linking as rapidly as possible.

- `VK_PIPELINE_CREATE_2_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`
  specifies that pipeline libraries should retain any information
  necessary to later perform an optimal link with
  `VK_PIPELINE_CREATE_2_LINK_TIME_OPTIMIZATION_BIT_EXT`.

- `VK_PIPELINE_CREATE_2_DESCRIPTOR_BUFFER_BIT_EXT` specifies that a
  pipeline will be used with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorbuffers"
  target="_blank" rel="noopener">descriptor buffers</a>, rather than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets"
  target="_blank" rel="noopener">descriptor sets</a>.

- `VK_PIPELINE_CREATE_2_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  specifies that the pipeline **may** be used with an attachment
  feedback loop including color attachments.

- `VK_PIPELINE_CREATE_2_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  specifies that the pipeline **may** be used with an attachment
  feedback loop including depth-stencil attachments.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT` specifies
  that the ray tracing pipeline **can** be used with acceleration
  structures which reference an opacity micromap array.

- `VK_PIPELINE_CREATE_2_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`
  specifies that the ray tracing pipeline **can** be used with
  acceleration structures which reference a displacement micromap array.

- `VK_PIPELINE_CREATE_2_NO_PROTECTED_ACCESS_BIT_EXT` specifies that the
  pipeline **must** not be bound to a protected command buffer.

- `VK_PIPELINE_CREATE_2_PROTECTED_ACCESS_ONLY_BIT_EXT` specifies that
  the pipeline **must** not be bound to an unprotected command buffer.

- `VK_PIPELINE_CREATE_2_ENABLE_LEGACY_DITHERING_BIT_EXT` specifies that
  the pipeline will be used in a render pass that is begun with
  `VK_RENDERING_ENABLE_LEGACY_DITHERING_BIT_EXT`.

It is valid to set both `VK_PIPELINE_CREATE_2_ALLOW_DERIVATIVES_BIT_KHR`
and `VK_PIPELINE_CREATE_2_DERIVATIVE_BIT_KHR`. This allows a pipeline to
be both a parent and possibly a child in a pipeline hierarchy. See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pipeline-derivatives"
target="_blank" rel="noopener">Pipeline Derivatives</a> for more
information.

When an implementation is looking up a pipeline in a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-cache"
target="_blank" rel="noopener">pipeline cache</a>, if that pipeline is
being created using linked libraries, implementations **should** always
return an equivalent pipeline created with
`VK_PIPELINE_CREATE_2_LINK_TIME_OPTIMIZATION_BIT_EXT` if available,
whether or not that bit was specified.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Using
<code>VK_PIPELINE_CREATE_2_LINK_TIME_OPTIMIZATION_BIT_EXT</code> (or
not) when linking pipeline libraries is intended as a performance
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

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkPipelineCreateFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCreateFlagBits2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
