# VK\_NV\_ray\_tracing(3) Manual Page

## Name

VK\_NV\_ray\_tracing - device extension



## [](#_registered_extension_number)Registered Extension Number

166

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK\_KHR\_get\_memory\_requirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_memory_requirements2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_1
- Interacts with VK\_EXT\_debug\_report
- Interacts with VK\_KHR\_get\_memory\_requirements2

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_ray\_tracing](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_ray_tracing.html)

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html) extension

## [](#_contact)Contact

- Eric Werness [\[GitHub\]ewerness-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_ray_tracing%5D%20%40ewerness-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_ray_tracing%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-11-20

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_ray_tracing`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_ray_tracing.txt)

**Contributors**

- Eric Werness, NVIDIA
- Ashwin Lele, NVIDIA
- Robert Stepinski, NVIDIA
- Nuno Subtil, NVIDIA
- Christoph Kubisch, NVIDIA
- Martin Stich, NVIDIA
- Daniel Koch, NVIDIA
- Jeff Bolz, NVIDIA
- Joshua Barczak, Intel
- Tobias Hector, AMD
- Henrik Rydgard, NVIDIA
- Pascal Gautron, NVIDIA

## [](#_description)Description

Rasterization has been the dominant method to produce interactive graphics, but increasing performance of graphics hardware has made ray tracing a viable option for interactive rendering. Being able to integrate ray tracing with traditional rasterization makes it easier for applications to incrementally add ray traced effects to existing applications or to do hybrid approaches with rasterization for primary visibility and ray tracing for secondary queries.

To enable ray tracing, this extension adds a few different categories of new functionality:

- Acceleration structure objects and build commands
- A new pipeline type with new shader domains
- An indirection table to link shader groups with acceleration structure items

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_NV_ray_tracing`

## [](#_new_object_types)New Object Types

- [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html)

## [](#_new_commands)New Commands

- [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindAccelerationStructureMemoryNV.html)
- [vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructureNV.html)
- [vkCmdCopyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureNV.html)
- [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysNV.html)
- [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html)
- [vkCompileDeferredNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCompileDeferredNV.html)
- [vkCreateAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateAccelerationStructureNV.html)
- [vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesNV.html)
- [vkDestroyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyAccelerationStructureNV.html)
- [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureHandleNV.html)
- [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)
- [vkGetRayTracingShaderGroupHandlesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupHandlesNV.html)

## [](#_new_structures)New Structures

- [VkAabbPositionsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAabbPositionsNV.html)
- [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoNV.html)
- [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)
- [VkAccelerationStructureInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceNV.html)
- [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)
- [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindAccelerationStructureMemoryInfoNV.html)
- [VkGeometryAABBNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryAABBNV.html)
- [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryDataNV.html)
- [VkGeometryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryNV.html)
- [VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTrianglesNV.html)
- [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)
- [VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoNV.html)
- [VkTransformMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)
- Extending [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html):
  
  - [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureNV.html)

If [VK\_KHR\_get\_memory\_requirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_memory_requirements2.html) or [Vulkan Version 1.1](#versions-1.1) is supported:

- [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2KHR.html)

## [](#_new_enums)New Enums

- [VkAccelerationStructureMemoryRequirementsTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsTypeNV.html)
- [VkAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTypeNV.html)
- [VkBuildAccelerationStructureFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsNV.html)
- [VkCopyAccelerationStructureModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeNV.html)
- [VkGeometryFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagBitsNV.html)
- [VkGeometryInstanceFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryInstanceFlagBitsNV.html)
- [VkGeometryTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeNV.html)
- [VkRayTracingShaderGroupTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkBuildAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagsNV.html)
- [VkGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagsNV.html)
- [VkGeometryInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryInstanceFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_RAY_TRACING_EXTENSION_NAME`
- `VK_NV_RAY_TRACING_SPEC_VERSION`
- `VK_SHADER_UNUSED_NV`
- Extending [VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTypeKHR.html):
  
  - `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV`
  - `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV`
- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_NV`
  - `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_NV`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV`
- Extending [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):
  
  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_NV`
  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV`
  - `VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_NV`
  - `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV`
  - `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV`
- Extending [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html):
  
  - `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV`
  - `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV`
- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html):
  
  - `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`
- Extending [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryFlagBitsKHR.html):
  
  - `VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_NV`
  - `VK_GEOMETRY_OPAQUE_BIT_NV`
- Extending [VkGeometryInstanceFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryInstanceFlagBitsKHR.html):
  
  - `VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_NV`
  - `VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_NV`
  - `VK_GEOMETRY_INSTANCE_TRIANGLE_CULL_DISABLE_BIT_NV`
  - `VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_NV`
- Extending [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html):
  
  - `VK_GEOMETRY_TYPE_AABBS_NV`
  - `VK_GEOMETRY_TYPE_TRIANGLES_NV`
- Extending [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html):
  
  - `VK_INDEX_TYPE_NONE_NV`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV`
- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html):
  
  - `VK_PIPELINE_BIND_POINT_RAY_TRACING_NV`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_NV`
  - `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_NV`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`
- Extending [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeKHR.html):
  
  - `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV`
  - `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`
  - `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV`
- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html):
  
  - `VK_SHADER_STAGE_ANY_HIT_BIT_NV`
  - `VK_SHADER_STAGE_CALLABLE_BIT_NV`
  - `VK_SHADER_STAGE_CLOSEST_HIT_BIT_NV`
  - `VK_SHADER_STAGE_INTERSECTION_BIT_NV`
  - `VK_SHADER_STAGE_MISS_BIT_NV`
  - `VK_SHADER_STAGE_RAYGEN_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_INFO_NV`
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV`
  - `VK_STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV`
  - `VK_STRUCTURE_TYPE_GEOMETRY_AABB_NV`
  - `VK_STRUCTURE_TYPE_GEOMETRY_NV`
  - `VK_STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV`

If [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html):
  
  - `VK_DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`LaunchIdNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-launchid)
- [`LaunchSizeNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-launchsize)
- [`WorldRayOriginNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-worldrayorigin)
- [`WorldRayDirectionNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-worldraydirection)
- [`ObjectRayOriginNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-objectrayorigin)
- [`ObjectRayDirectionNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-objectraydirection)
- [`RayTminNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-raytmin)
- [`RayTmaxNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-raytmax)
- [`InstanceCustomIndexNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-instancecustomindex)
- [`InstanceId`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-instanceid)
- [`ObjectToWorldNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-objecttoworld)
- [`WorldToObjectNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-worldtoobject)
- [`HitTNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitt)
- [`HitKindNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitkind)
- [`IncomingRayFlagsNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-incomingrayflags)
- (modified)`PrimitiveId`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`RayTracingNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTracingNV)

## [](#_issues)Issues

1\) Are there issues?

**RESOLVED**: Yes.

## [](#_sample_code)Sample Code

Example ray generation GLSL shader

```c
#version 450 core
#extension GL_NV_ray_tracing : require
layout(set = 0, binding = 0, rgba8) uniform image2D image;
layout(set = 0, binding = 1) uniform accelerationStructureNV as;
layout(location = 0) rayPayloadNV float payload;

void main()
{
   vec4 col = vec4(0, 0, 0, 1);

   vec3 origin = vec3(float(gl_LaunchIDNV.x)/float(gl_LaunchSizeNV.x), float(gl_LaunchIDNV.y)/float(gl_LaunchSizeNV.y), 1.0);
   vec3 dir = vec3(0.0, 0.0, -1.0);

   traceNV(as, 0, 0xff, 0, 1, 0, origin, 0.0, dir, 1000.0, 0);

   col.y = payload;

   imageStore(image, ivec2(gl_LaunchIDNV.xy), col);
}
```

## [](#_version_history)Version History

- Revision 1, 2018-09-11 (Robert Stepinski, Nuno Subtil, Eric Werness)
  
  - Internal revisions
- Revision 2, 2018-10-19 (Eric Werness)
  
  - rename to VK\_NV\_ray\_tracing, add support for callables.
  - too many updates to list
- Revision 3, 2018-11-20 (Daniel Koch)
  
  - update to use InstanceId instead of InstanceIndex as implemented.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_ray_tracing)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0