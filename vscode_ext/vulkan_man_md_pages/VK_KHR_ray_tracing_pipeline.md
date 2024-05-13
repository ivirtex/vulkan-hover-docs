# VK_KHR_ray_tracing_pipeline(3) Manual Page

## Name

VK_KHR_ray_tracing_pipeline - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

348

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_spirv_1_4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_spirv_1_4.html)  
and  
[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_ray_tracing](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_ray_tracing.html)

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_ray_tracing_pipeline%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_ray_tracing_pipeline%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-11-12

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_ray_tracing`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_tracing.txt)

- This extension interacts with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2"
  target="_blank" rel="noopener">Vulkan 1.2</a> and
  [`VK_KHR_vulkan_memory_model`](VK_KHR_vulkan_memory_model.html),
  adding the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shader-call-related"
  target="_blank" rel="noopener">shader-call-related</a> relation of
  invocations, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shader-call-order"
  target="_blank" rel="noopener">shader-call-order</a> partial order of
  dynamic instances of instructions, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-shadercall"
  target="_blank" rel="noopener"><code>ShaderCallKHR</code></a> scope.

- This extension interacts with
  [`VK_KHR_pipeline_library`](VK_KHR_pipeline_library.html), enabling
  pipeline libraries to be used with ray tracing pipelines and enabling
  usage of
  [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html).

**Contributors**  
- Matthäus Chajdas, AMD

- Greg Grebe, AMD

- Nicolai Hähnle, AMD

- Tobias Hector, AMD

- Dave Oldcorn, AMD

- Skyler Saleh, AMD

- Mathieu Robart, Arm

- Marius Bjorge, Arm

- Tom Olson, Arm

- Sebastian Tafuri, EA

- Henrik Rydgard, Embark

- Juan Cañada, Epic Games

- Patrick Kelly, Epic Games

- Yuriy O’Donnell, Epic Games

- Michael Doggett, Facebook/Oculus

- Andrew Garrard, Imagination

- Don Scorgie, Imagination

- Dae Kim, Imagination

- Joshua Barczak, Intel

- Slawek Grajewski, Intel

- Jeff Bolz, NVIDIA

- Pascal Gautron, NVIDIA

- Daniel Koch, NVIDIA

- Christoph Kubisch, NVIDIA

- Ashwin Lele, NVIDIA

- Robert Stepinski, NVIDIA

- Martin Stich, NVIDIA

- Nuno Subtil, NVIDIA

- Eric Werness, NVIDIA

- Jon Leech, Khronos

- Jeroen van Schijndel, OTOY

- Juul Joosten, OTOY

- Alex Bourd, Qualcomm

- Roman Larionov, Qualcomm

- David McAllister, Qualcomm

- Spencer Fricke, Samsung

- Lewis Gordon, Samsung

- Ralph Potter, Samsung

- Jasper Bekkers, Traverse Research

- Jesse Barker, Unity

- Baldur Karlsson, Valve

## <a href="#_description" class="anchor"></a>Description

Rasterization has been the dominant method to produce interactive
graphics, but increasing performance of graphics hardware has made ray
tracing a viable option for interactive rendering. Being able to
integrate ray tracing with traditional rasterization makes it easier for
applications to incrementally add ray traced effects to existing
applications or to do hybrid approaches with rasterization for primary
visibility and ray tracing for secondary queries.

To enable ray tracing, this extension adds a few different categories of
new functionality:

- A new ray tracing pipeline type with new shader domains: ray
  generation, intersection, any-hit, closest hit, miss, and callable

- A shader binding indirection table to link shader groups with
  acceleration structure items

- Ray tracing commands which initiate the ray pipeline traversal and
  invocation of the various new shader domains depending on which
  traversal conditions are met

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_KHR_ray_tracing`

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html)

- [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html)

- [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html)

- [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html)

- [vkGetRayTracingCaptureReplayShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingCaptureReplayShaderGroupHandlesKHR.html)

- [vkGetRayTracingShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingShaderGroupHandlesKHR.html)

- [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)

- [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)

- [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)

- [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)

- [VkTraceRaysIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTraceRaysIndirectCommandKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupTypeKHR.html)

- [VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderGroupShaderKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_RAY_TRACING_PIPELINE_EXTENSION_NAME`

- `VK_KHR_RAY_TRACING_PIPELINE_SPEC_VERSION`

- `VK_SHADER_UNUSED_KHR`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR`

- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html):

  - `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`

  - `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

  - `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`

  - `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html):

  - `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html):

  - `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`

  - `VK_SHADER_STAGE_CALLABLE_BIT_KHR`

  - `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`

  - `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`

  - `VK_SHADER_STAGE_MISS_BIT_KHR`

  - `VK_SHADER_STAGE_RAYGEN_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_INTERFACE_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_KHR`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-launchid"
  target="_blank" rel="noopener"><code>LaunchIdKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-launchsize"
  target="_blank" rel="noopener"><code>LaunchSizeKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-worldrayorigin"
  target="_blank" rel="noopener"><code>WorldRayOriginKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-worldraydirection"
  target="_blank" rel="noopener"><code>WorldRayDirectionKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-objectrayorigin"
  target="_blank" rel="noopener"><code>ObjectRayOriginKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-objectraydirection"
  target="_blank" rel="noopener"><code>ObjectRayDirectionKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-raytmin"
  target="_blank" rel="noopener"><code>RayTminKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-raytmax"
  target="_blank" rel="noopener"><code>RayTmaxKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-instancecustomindex"
  target="_blank" rel="noopener"><code>InstanceCustomIndexKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-instanceid"
  target="_blank" rel="noopener"><code>InstanceId</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-objecttoworld"
  target="_blank" rel="noopener"><code>ObjectToWorldKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-worldtoobject"
  target="_blank" rel="noopener"><code>WorldToObjectKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-hitkind"
  target="_blank" rel="noopener"><code>HitKindKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-incomingrayflags"
  target="_blank" rel="noopener"><code>IncomingRayFlagsKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-raygeometryindex"
  target="_blank" rel="noopener"><code>RayGeometryIndexKHR</code></a>

- (modified)`PrimitiveId`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayTracingKHR"
  target="_blank" rel="noopener"><code>RayTracingKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayTraversalPrimitiveCullingKHR"
  target="_blank"
  rel="noopener"><code>RayTraversalPrimitiveCullingKHR</code></a>

## <a href="#_issues" class="anchor"></a>Issues

\(1\) How does this extension differ from VK_NV_ray_tracing?

**DISCUSSION**:

The following is a summary of the main functional differences between
VK_KHR_ray_tracing_pipeline and VK_NV_ray_tracing:

- added support for indirect ray tracing
  ([vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html))

- uses SPV_KHR_ray_tracing instead of SPV_NV_ray_tracing

  - refer to KHR SPIR-V enums instead of NV SPIR-V enums (which are
    functionally equivalent and aliased to the same values).

  - added `RayGeometryIndexKHR` built-in

- removed vkCompileDeferredNV compilation functionality and replaced
  with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations"
  target="_blank" rel="noopener">deferred host operations</a>
  interactions for ray tracing

- added
  [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)
  structure

- extended
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)
  structure

  - renamed `maxRecursionDepth` to `maxRayRecursionDepth` and it has a
    minimum of 1 instead of 31

  - require `shaderGroupHandleSize` to be 32 bytes

  - added `maxRayDispatchInvocationCount`, `shaderGroupHandleAlignment`
    and `maxRayHitAttributeSize`

- reworked geometry structures so they could be better shared between
  device, host, and indirect builds

- changed SBT parameters to a structure and added size
  ([VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html))

- add parameter for requesting memory requirements for host and/or
  device build

- added <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-library"
  target="_blank" rel="noopener">pipeline library</a> support for ray
  tracing

- added <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-traversal-watertight"
  target="_blank" rel="noopener">watertightness guarantees</a>

- added no-null-shader pipeline flags
  (`VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_*_SHADERS_BIT_KHR`)

- added <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-shader-call"
  target="_blank" rel="noopener">memory model interactions</a> with ray
  tracing and define how subgroups work and can be repacked

\(2\) Can you give a more detailed comparison of differences and
similarities between VK_NV_ray_tracing and VK_KHR_ray_tracing_pipeline?

**DISCUSSION**:

The following is a more detailed comparison of which commands,
structures, and enums are aliased, changed, or removed.

- Aliased functionality — enums, structures, and commands that are
  considered equivalent:

  - [VkRayTracingShaderGroupTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupTypeNV.html)
    ↔
    [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupTypeKHR.html)

  - [vkGetRayTracingShaderGroupHandlesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingShaderGroupHandlesNV.html)
    ↔
    [vkGetRayTracingShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingShaderGroupHandlesKHR.html)

- Changed enums, structures, and commands:

  - [VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoNV.html)
    →
    [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)
    (added `pShaderGroupCaptureReplayHandle`)

  - [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)
    →
    [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)
    (changed type of `pGroups`, added `libraries`, `pLibraryInterface`,
    and `pDynamicState`)

  - [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)
    → VkPhysicalDeviceRayTracingPropertiesKHR (renamed
    `maxTriangleCount` to `maxPrimitiveCount`, added
    `shaderGroupHandleCaptureReplaySize`)

  - [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysNV.html) →
    [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html) (params to struct)

  - [vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesNV.html)
    →
    [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html)
    (different struct, changed functionality)

- Added enums, structures and commands:

  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
    `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`,
    `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`,
    `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`,
    `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`,
    `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR` to
    [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html)

  - [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)
    structure

  - [VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html) and
    [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html)
    unions

  - [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
    struct

  - [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)
    struct

  - [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
    struct

  - [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html) command
    and
    [VkTraceRaysIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTraceRaysIndirectCommandKHR.html)
    struct

  - [vkGetRayTracingCaptureReplayShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingCaptureReplayShaderGroupHandlesKHR.html)
    (shader group capture/replay)

  - [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html)
    and
    [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html)
    commands for stack size control

- Functionality removed:

  - `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV`

  - [vkCompileDeferredNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCompileDeferredNV.html) command (replaced
    with
    [`VK_KHR_deferred_host_operations`](VK_KHR_deferred_host_operations.html))

\(3\) What are the changes between the public provisional
(VK_KHR_ray_tracing v8) release and the internal provisional
(VK_KHR_ray_tracing v9) release?

- Require Vulkan 1.1 and SPIR-V 1.4

- Added interactions with Vulkan 1.2 and
  [`VK_KHR_vulkan_memory_model`](VK_KHR_vulkan_memory_model.html)

- added creation time capture and replay flags

  - added
    `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
    to [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html)

- replace `VkStridedBufferRegionKHR` with
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  and change [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html),
  [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html), to take
  these for the shader binding table and use device addresses instead of
  buffers.

- require the shader binding table buffers to have the
  `VK_BUFFER_USAGE_RAY_TRACING_BIT_KHR` set

- make [`VK_KHR_pipeline_library`](VK_KHR_pipeline_library.html) an
  interaction instead of required extension

- rename the `libraries` member of
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)
  to `pLibraryInfo` and make it a pointer

- make
  [`VK_KHR_deferred_host_operations`](VK_KHR_deferred_host_operations.html)
  an interaction instead of a required extension (later went back on
  this)

- added explicit stack size management for ray tracing pipelines

  - removed the `maxCallableSize` member of
    [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)

  - added the `pDynamicState` member to
    [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)

  - added `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` dynamic
    state for ray tracing pipelines

  - added
    [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html)
    and
    [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html)
    commands

  - added [VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderGroupShaderKHR.html) enum

- Added `maxRayDispatchInvocationCount` limit to
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)

- Added `shaderGroupHandleAlignment` property to
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)

- Added `maxRayHitAttributeSize` property to
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)

- Clarify deferred host ops for pipeline creation

  - [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) is now a
    top-level parameter for
    [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html)

  - removed `VkDeferredOperationInfoKHR` structure

  - change deferred host creation/return parameter behavior such that
    the implementation can modify such parameters until the deferred
    host operation completes

  - [`VK_KHR_deferred_host_operations`](VK_KHR_deferred_host_operations.html)
    is required again

\(4\) What are the changes between the internal provisional
(VK_KHR_ray_tracing v9) release and the final
(VK_KHR_acceleration_structure v11 / VK_KHR_ray_tracing_pipeline v1)
release?

- refactor VK_KHR_ray_tracing into 3 extensions, enabling implementation
  flexibility and decoupling ray query support from ray pipelines:

  - [`VK_KHR_acceleration_structure`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)
    (for acceleration structure operations)

  - [`VK_KHR_ray_tracing_pipeline`](VK_KHR_ray_tracing_pipeline.html)
    (for ray tracing pipeline and shader stages)

  - [`VK_KHR_ray_query`](VK_KHR_ray_query.html) (for ray queries in
    existing shader stages)

- Require `Volatile` for the following builtins in the ray generation,
  closest hit, miss, intersection, and callable shader stages:

  - `SubgroupSize`, `SubgroupLocalInvocationId`, `SubgroupEqMask`,
    `SubgroupGeMask`, `SubgroupGtMask`, `SubgroupLeMask`,
    `SubgroupLtMask`

  - `SMIDNV`, `WarpIDNV`

- clarify buffer usage flags for ray tracing

  - `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` is added as an alias
    of `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` and is required on shader
    binding table buffers

  - `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` is used in
    [`VK_KHR_acceleration_structure`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)
    for `scratchData`

- rename `maxRecursionDepth` to `maxRayPipelineRecursionDepth` (pipeline
  creation) and `maxRayRecursionDepth` (limit) to reduce confusion

- Add queryable `maxRayHitAttributeSize` limit and rename members of
  [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)
  to `maxPipelineRayPayloadSize` and `maxPipelineRayHitAttributeSize`
  for clarity

- Update SPIRV capabilities to use `RayTracingKHR`

- extension is no longer provisional

- define synchronization requirements for indirect trace rays and
  indirect buffer

\(5\) This extension adds gl_InstanceID for the intersection, any-hit,
and closest hit shaders, but in KHR_vulkan_glsl, gl_InstanceID is
replaced with gl_InstanceIndex. Which should be used for Vulkan in this
extension?

**RESOLVED**: This extension uses gl_InstanceID and maps it to
`InstanceId` in SPIR-V. It is acknowledged that this is different than
other shader stages in Vulkan. There are two main reasons for the
difference here:

- symmetry with gl_PrimitiveID which is also available in these shaders

- there is no “baseInstance” relevant for these shaders, and so ID makes
  it more obvious that this is zero-based.

\(6\) Why is [`VK_KHR_pipeline_library`](VK_KHR_pipeline_library.html)
an interaction instead of a required dependency, particularly when the
“Feature Requirements” section says it is required to be supported
anyhow?

**RESOLVED**: If
[`VK_KHR_pipeline_library`](VK_KHR_pipeline_library.html) were a
required extension dependency, then every application would need to
enable the extension whether or not they actually want to use the
pipeline library functionality. Developers found this to be annoying and
unfriendly behavior. We do wish to require all **implementations** to
support it though, and thus it is listed in the feature requirements
section.

## <a href="#_sample_code" class="anchor"></a>Sample Code

Example ray generation GLSL shader

``` c
#version 450 core
#extension GL_EXT_ray_tracing : require
layout(set = 0, binding = 0, rgba8) uniform image2D image;
layout(set = 0, binding = 1) uniform accelerationStructureEXT as;
layout(location = 0) rayPayloadEXT float payload;

void main()
{
   vec4 col = vec4(0, 0, 0, 1);

   vec3 origin = vec3(float(gl_LaunchIDEXT.x)/float(gl_LaunchSizeEXT.x), float(gl_LaunchIDEXT.y)/float(gl_LaunchSizeEXT.y), 1.0);
   vec3 dir = vec3(0.0, 0.0, -1.0);

   traceRayEXT(as, 0, 0xff, 0, 1, 0, origin, 0.0, dir, 1000.0, 0);

   col.y = payload;

   imageStore(image, ivec2(gl_LaunchIDEXT.xy), col);
}
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-11-12 (Mathieu Robart, Daniel Koch, Eric Werness,
  Tobias Hector)

  - Decomposition of the specification, from VK_KHR_ray_tracing to
    VK_KHR_ray_tracing_pipeline (#1918,!3912)

  - require certain subgroup and sm_shader_builtin shader builtins to be
    decorated as volatile in the ray generation, closest hit, miss,
    intersection, and callable stages (#1924,!3903,!3954)

  - clarify buffer usage flags for ray tracing (#2181,!3939)

  - rename maxRecursionDepth to maxRayPipelineRecursionDepth and
    maxRayRecursionDepth (#2203,!3937)

  - add queryable maxRayHitAttributeSize and rename members of
    VkRayTracingPipelineInterfaceCreateInfoKHR (#2102,!3966)

  - update to use `RayTracingKHR` SPIR-V capability

  - add VUs for matching hit group type against geometry type
    (#2245,!3994)

  - require `RayTMaxKHR` be volatile in intersection shaders
    (#2268,!4030)

  - add numerical limits for ray parameters (#2235,!3960)

  - fix SBT indexing rules for device addresses (#2308,!4079)

  - relax formula for ray intersection candidate determination
    (#2322,!4080)

  - add more details on `ShaderRecordBufferKHR` variables (#2230,!4083)

  - clarify valid bits for `InstanceCustomIndexKHR` (GLSL/GLSL#19,!4128)

  - allow at most one `IncomingRayPayloadKHR`,
    `IncomingCallableDataKHR`, and `HitAttributeKHR` (!4129)

  - add minimum for maxShaderGroupStride (#2353,!4131)

  - require VK_KHR_pipeline_library extension to be supported
    (#2348,!4135)

  - clarify meaning of 'geometry index' (#2272,!4137)

  - restrict traces to TLAS (#2239,!4141)

  - add note about maxPipelineRayPayloadSize (#2383,!4172)

  - do not require raygen shader in pipeline libraries (!4185)

  - define sync for indirect trace rays and indirect buffer
    (#2407,!4208)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_ray_tracing_pipeline"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
