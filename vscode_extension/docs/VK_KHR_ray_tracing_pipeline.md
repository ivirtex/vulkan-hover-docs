# VK\_KHR\_ray\_tracing\_pipeline(3) Manual Page

## Name

VK\_KHR\_ray\_tracing\_pipeline - device extension



## [](#_registered_extension_number)Registered Extension Number

348

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_spirv\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_spirv_1_4.html)  
     or  
     [Vulkan Version 1.2](#versions-1.2)  
and  
[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_ray\_query

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_ray\_tracing](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_ray_tracing.html)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_ray_tracing_pipeline%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_ray_tracing_pipeline%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-11-12

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_ray_tracing`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_tracing.txt)
- This extension interacts with [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2) and `VK_KHR_vulkan_memory_model`, adding the [shader-call-related](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shader-call-related) relation of invocations, [shader-call-order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shader-call-order) partial order of dynamic instances of instructions, and the [`ShaderCallKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-scope-shadercall) scope.
- This extension interacts with `VK_KHR_pipeline_library`, enabling pipeline libraries to be used with ray tracing pipelines and enabling usage of [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html).

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

## [](#_description)Description

Rasterization has been the dominant method to produce interactive graphics, but increasing performance of graphics hardware has made ray tracing a viable option for interactive rendering. Being able to integrate ray tracing with traditional rasterization makes it easier for applications to incrementally add ray traced effects to existing applications or to do hybrid approaches with rasterization for primary visibility and ray tracing for secondary queries.

To enable ray tracing, this extension adds a few different categories of new functionality:

- A new ray tracing pipeline type with new shader domains: ray generation, intersection, any-hit, closest hit, miss, and callable
- A shader binding indirection table to link shader groups with acceleration structure items
- Ray tracing commands which initiate the ray pipeline traversal and invocation of the various new shader domains depending on which traversal conditions are met

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_KHR_ray_tracing`

## [](#_new_commands)New Commands

- [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html)
- [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html)
- [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html)
- [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesKHR.html)
- [vkGetRayTracingCaptureReplayShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingCaptureReplayShaderGroupHandlesKHR.html)
- [vkGetRayTracingShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupHandlesKHR.html)
- [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html)

## [](#_new_structures)New Structures

- [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)
- [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)
- [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)
- [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html)
- [VkTraceRaysIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTraceRaysIndirectCommandKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeKHR.html)
- [VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderGroupShaderKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_RAY_TRACING_PIPELINE_EXTENSION_NAME`
- `VK_KHR_RAY_TRACING_PIPELINE_SPEC_VERSION`
- `VK_SHADER_UNUSED_KHR`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR`
- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html):
  
  - `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`
  - `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
  - `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`
  - `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_SKIP_BUILT_IN_PRIMITIVES_BIT_KHR`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html):
  
  - `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`
  - `VK_SHADER_STAGE_CALLABLE_BIT_KHR`
  - `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`
  - `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`
  - `VK_SHADER_STAGE_MISS_BIT_KHR`
  - `VK_SHADER_STAGE_RAYGEN_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_INTERFACE_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_KHR`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`LaunchIdKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-launchid)
- [`LaunchSizeKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-launchsize)
- [`WorldRayOriginKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-worldrayorigin)
- [`WorldRayDirectionKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-worldraydirection)
- [`ObjectRayOriginKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-objectrayorigin)
- [`ObjectRayDirectionKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-objectraydirection)
- [`RayTminKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-raytmin)
- [`RayTmaxKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-raytmax)
- [`InstanceCustomIndexKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-instancecustomindex)
- [`InstanceId`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-instanceid)
- [`ObjectToWorldKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-objecttoworld)
- [`WorldToObjectKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-worldtoobject)
- [`HitKindKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitkind)
- [`IncomingRayFlagsKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-incomingrayflags)
- [`RayGeometryIndexKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-raygeometryindex)
- (modified)`PrimitiveId`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`RayTracingKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTracingKHR)
- [`RayTraversalPrimitiveCullingKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTraversalPrimitiveCullingKHR)

## [](#_issues)Issues

(1) How does this extension differ from VK\_NV\_ray\_tracing?

**DISCUSSION**:

The following is a summary of the main functional differences between VK\_KHR\_ray\_tracing\_pipeline and VK\_NV\_ray\_tracing:

- added support for indirect ray tracing ([vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html))
- uses SPV\_KHR\_ray\_tracing instead of SPV\_NV\_ray\_tracing
  
  - refer to KHR SPIR-V enums instead of NV SPIR-V enums (which are functionally equivalent and aliased to the same values).
  - added `RayGeometryIndexKHR` built-in
- removed vkCompileDeferredNV compilation functionality and replaced with [deferred host operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations) interactions for ray tracing
- added [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html) structure
- extended [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html) structure
  
  - renamed `maxRecursionDepth` to `maxRayRecursionDepth` and it has a minimum of 1 instead of 31
  - require `shaderGroupHandleSize` to be 32 bytes
  - added `maxRayDispatchInvocationCount`, `shaderGroupHandleAlignment` and `maxRayHitAttributeSize`
- reworked geometry structures so they could be better shared between device, host, and indirect builds
- changed SBT parameters to a structure and added size ([VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html))
- add parameter for requesting memory requirements for host and/or device build
- added [pipeline library](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-library) support for ray tracing
- added [watertightness guarantees](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-traversal-watertight)
- added no-null-shader pipeline flags (`VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_*_SHADERS_BIT_KHR`)
- added [memory model interactions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing-shader-call) with ray tracing and define how subgroups work and can be repacked

(2) Can you give a more detailed comparison of differences and similarities between VK\_NV\_ray\_tracing and VK\_KHR\_ray\_tracing\_pipeline?

**DISCUSSION**:

The following is a more detailed comparison of which commands, structures, and enums are aliased, changed, or removed.

- Aliased functionality — enums, structures, and commands that are considered equivalent:
  
  - [VkRayTracingShaderGroupTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeNV.html) ↔ [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeKHR.html)
  - [vkGetRayTracingShaderGroupHandlesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupHandlesNV.html) ↔ [vkGetRayTracingShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupHandlesKHR.html)
- Changed enums, structures, and commands:
  
  - [VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoNV.html) → [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html) (added `pShaderGroupCaptureReplayHandle`)
  - [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html) → [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html) (changed type of `pGroups`, added `libraries`, `pLibraryInterface`, and `pDynamicState`)
  - [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html) → [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html) (renamed `maxTriangleCount` to `maxPrimitiveCount`, added `shaderGroupHandleCaptureReplaySize`)
  - [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysNV.html) → [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html) (params to struct)
  - [vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesNV.html) → [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesKHR.html) (different struct, changed functionality)
- Added enums, structures and commands:
  
  - `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR` `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`, `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`, `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`, `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`, `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR` to [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html)
  - [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html) structure
  - [VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressKHR.html) and [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) unions
  - [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html) struct
  - [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html) struct
  - [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) struct
  - [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html) command and [VkTraceRaysIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTraceRaysIndirectCommandKHR.html) struct
  - [vkGetRayTracingCaptureReplayShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingCaptureReplayShaderGroupHandlesKHR.html) (shader group capture/replay)
  - [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html) and [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html) commands for stack size control
- Functionality removed:
  
  - `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV`
  - [vkCompileDeferredNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCompileDeferredNV.html) command (replaced with `VK_KHR_deferred_host_operations`)

(3) What are the changes between the public provisional (VK\_KHR\_ray\_tracing v8) release and the internal provisional (VK\_KHR\_ray\_tracing v9) release?

- Require Vulkan 1.1 and SPIR-V 1.4
- Added interactions with Vulkan 1.2 and `VK_KHR_vulkan_memory_model`
- added creation time capture and replay flags
  
  - added `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR` to [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html)
- replace `VkStridedBufferRegionKHR` with [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) and change [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html), [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html), to take these for the shader binding table and use device addresses instead of buffers.
- require the shader binding table buffers to have the `VK_BUFFER_USAGE_RAY_TRACING_BIT_KHR` set
- make `VK_KHR_pipeline_library` an interaction instead of required extension
- rename the `libraries` member of [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html) to `pLibraryInfo` and make it a pointer
- make `VK_KHR_deferred_host_operations` an interaction instead of a required extension (later went back on this)
- added explicit stack size management for ray tracing pipelines
  
  - removed the `maxCallableSize` member of [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)
  - added the `pDynamicState` member to [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)
  - added `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` dynamic state for ray tracing pipelines
  - added [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html) and [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html) commands
  - added [VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderGroupShaderKHR.html) enum
- Added `maxRayDispatchInvocationCount` limit to [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)
- Added `shaderGroupHandleAlignment` property to [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)
- Added `maxRayHitAttributeSize` property to [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)
- Clarify deferred host ops for pipeline creation
  
  - [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) is now a top-level parameter for [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesKHR.html)
  - removed `VkDeferredOperationInfoKHR` structure
  - change deferred host creation/return parameter behavior such that the implementation can modify such parameters until the deferred host operation completes
  - `VK_KHR_deferred_host_operations` is required again

(4) What are the changes between the internal provisional (VK\_KHR\_ray\_tracing v9) release and the final (VK\_KHR\_acceleration\_structure v11 / VK\_KHR\_ray\_tracing\_pipeline v1) release?

- refactor VK\_KHR\_ray\_tracing into 3 extensions, enabling implementation flexibility and decoupling ray query support from ray pipelines:
  
  - `VK_KHR_acceleration_structure` (for acceleration structure operations)
  - `VK_KHR_ray_tracing_pipeline` (for ray tracing pipeline and shader stages)
  - `VK_KHR_ray_query` (for ray queries in existing shader stages)
- Require `Volatile` for the following builtins in the ray generation, closest hit, miss, intersection, and callable shader stages:
  
  - `SubgroupSize`, `SubgroupLocalInvocationId`, `SubgroupEqMask`, `SubgroupGeMask`, `SubgroupGtMask`, `SubgroupLeMask`, `SubgroupLtMask`
  - `SMIDNV`, `WarpIDNV`
- clarify buffer usage flags for ray tracing
  
  - `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` is added as an alias of `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` and is required on shader binding table buffers
  - `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` is used in `VK_KHR_acceleration_structure` for `scratchData`
- rename `maxRecursionDepth` to `maxRayPipelineRecursionDepth` (pipeline creation) and `maxRayRecursionDepth` (limit) to reduce confusion
- Add queryable `maxRayHitAttributeSize` limit and rename members of [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html) to `maxPipelineRayPayloadSize` and `maxPipelineRayHitAttributeSize` for clarity
- Update SPIRV capabilities to use `RayTracingKHR`
- extension is no longer provisional
- define synchronization requirements for indirect trace rays and indirect buffer

(5) This extension adds gl\_InstanceID for the intersection, any-hit, and closest hit shaders, but in KHR\_vulkan\_glsl, gl\_InstanceID is replaced with gl\_InstanceIndex. Which should be used for Vulkan in this extension?

**RESOLVED**: This extension uses gl\_InstanceID and maps it to `InstanceId` in SPIR-V. It is acknowledged that this is different than other shader stages in Vulkan. There are two main reasons for the difference here:

- symmetry with gl\_PrimitiveID which is also available in these shaders
- there is no “baseInstance” relevant for these shaders, and so ID makes it more obvious that this is zero-based.

(6) Why is `VK_KHR_pipeline_library` an interaction instead of a required dependency, particularly when the “Feature Requirements” section says it is required to be supported anyhow?

**RESOLVED**: If the `VK_KHR_pipeline_library` extension were a required dependency, then every application would need to enable the extension whether or not they actually want to use the pipeline library functionality. Developers found this to be annoying and unfriendly behavior. We do wish to require all **implementations** to support it though, and thus it is listed in the feature requirements section.

## [](#_sample_code)Sample Code

Example ray generation GLSL shader

```c
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

## [](#_version_history)Version History

- Revision 1, 2020-11-12 (Mathieu Robart, Daniel Koch, Eric Werness, Tobias Hector)
  
  - Decomposition of the specification, from VK\_KHR\_ray\_tracing to VK\_KHR\_ray\_tracing\_pipeline (#1918,!3912)
  - require certain subgroup and sm\_shader\_builtin shader builtins to be decorated as volatile in the ray generation, closest hit, miss, intersection, and callable stages (#1924,!3903,!3954)
  - clarify buffer usage flags for ray tracing (#2181,!3939)
  - rename maxRecursionDepth to maxRayPipelineRecursionDepth and maxRayRecursionDepth (#2203,!3937)
  - add queryable maxRayHitAttributeSize and rename members of VkRayTracingPipelineInterfaceCreateInfoKHR (#2102,!3966)
  - update to use `RayTracingKHR` SPIR-V capability
  - add VUs for matching hit group type against geometry type (#2245,!3994)
  - require `RayTMaxKHR` be volatile in intersection shaders (#2268,!4030)
  - add numerical limits for ray parameters (#2235,!3960)
  - fix SBT indexing rules for device addresses (#2308,!4079)
  - relax formula for ray intersection candidate determination (#2322,!4080)
  - add more details on `ShaderRecordBufferKHR` variables (#2230,!4083)
  - clarify valid bits for `InstanceCustomIndexKHR` (GLSL/GLSL#19,!4128)
  - allow at most one `IncomingRayPayloadKHR`, `IncomingCallableDataKHR`, and `HitAttributeKHR` (!4129)
  - add minimum for maxShaderGroupStride (#2353,!4131)
  - require VK\_KHR\_pipeline\_library extension to be supported (#2348,!4135)
  - clarify meaning of 'geometry index' (#2272,!4137)
  - restrict traces to TLAS (#2239,!4141)
  - add note about maxPipelineRayPayloadSize (#2383,!4172)
  - do not require raygen shader in pipeline libraries (!4185)
  - define sync for indirect trace rays and indirect buffer (#2407,!4208)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_ray_tracing_pipeline).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0