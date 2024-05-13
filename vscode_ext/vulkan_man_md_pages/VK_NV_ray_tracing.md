# VK_NV_ray_tracing(3) Manual Page

## Name

VK_NV_ray_tracing - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

166

## <a href="#_revision" class="anchor"></a>Revision

3

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
    
[VK_KHR_get_memory_requirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_memory_requirements2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_1

- Interacts with VK_EXT_debug_report

- Interacts with VK_KHR_get_memory_requirements2

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_ray_tracing](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_ray_tracing.html)

## <a href="#_contact" class="anchor"></a>Contact

- Eric Werness <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_ray_tracing%5D%20@ewerness-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_ray_tracing%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ewerness-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-11-20

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_ray_tracing`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_ray_tracing.txt)

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

- Acceleration structure objects and build commands

- A new pipeline type with new shader domains

- An indirection table to link shader groups with acceleration structure
  items

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_NV_ray_tracing`

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindAccelerationStructureMemoryNV.html)

- [vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructureNV.html)

- [vkCmdCopyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureNV.html)

- [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysNV.html)

- [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html)

- [vkCompileDeferredNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCompileDeferredNV.html)

- [vkCreateAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAccelerationStructureNV.html)

- [vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesNV.html)

- [vkDestroyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyAccelerationStructureNV.html)

- [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureHandleNV.html)

- [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)

- [vkGetRayTracingShaderGroupHandlesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRayTracingShaderGroupHandlesNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkAabbPositionsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAabbPositionsNV.html)

- [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoNV.html)

- [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)

- [VkAccelerationStructureInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceNV.html)

- [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)

- [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindAccelerationStructureMemoryInfoNV.html)

- [VkGeometryAABBNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryAABBNV.html)

- [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryDataNV.html)

- [VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html)

- [VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTrianglesNV.html)

- [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)

- [VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoNV.html)

- [VkTransformMatrixNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTransformMatrixNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)

- Extending [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html):

  - [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureNV.html)

If
[VK_KHR_get_memory_requirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_memory_requirements2.html)
or [Version 1.1](#versions-1.1) is supported:

- [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2KHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkAccelerationStructureMemoryRequirementsTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsTypeNV.html)

- [VkAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeNV.html)

- [VkBuildAccelerationStructureFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsNV.html)

- [VkCopyAccelerationStructureModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeNV.html)

- [VkGeometryFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagBitsNV.html)

- [VkGeometryInstanceFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagBitsNV.html)

- [VkGeometryTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeNV.html)

- [VkRayTracingShaderGroupTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupTypeNV.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkBuildAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagsNV.html)

- [VkGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagsNV.html)

- [VkGeometryInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagsNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_RAY_TRACING_EXTENSION_NAME`

- `VK_NV_RAY_TRACING_SPEC_VERSION`

- `VK_SHADER_UNUSED_NV`

- Extending
  [VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html):

  - `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV`

  - `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV`

- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html):

  - `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_NV`

  - `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_NV`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV`

- Extending
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):

  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_NV`

  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV`

  - `VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_NV`

  - `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV`

  - `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV`

- Extending
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html):

  - `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV`

  - `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV`

- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html):

  - `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`

- Extending [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagBitsKHR.html):

  - `VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_NV`

  - `VK_GEOMETRY_OPAQUE_BIT_NV`

- Extending
  [VkGeometryInstanceFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagBitsKHR.html):

  - `VK_GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_NV`

  - `VK_GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_NV`

  - `VK_GEOMETRY_INSTANCE_TRIANGLE_CULL_DISABLE_BIT_NV`

  - `VK_GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_NV`

- Extending [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeKHR.html):

  - `VK_GEOMETRY_TYPE_AABBS_NV`

  - `VK_GEOMETRY_TYPE_TRIANGLES_NV`

- Extending [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html):

  - `VK_INDEX_TYPE_NONE_NV`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV`

- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html):

  - `VK_PIPELINE_BIND_POINT_RAY_TRACING_NV`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV`

- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html):

  - `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_NV`

  - `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_NV`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`

- Extending
  [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupTypeKHR.html):

  - `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV`

  - `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`

  - `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV`

- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html):

  - `VK_SHADER_STAGE_ANY_HIT_BIT_NV`

  - `VK_SHADER_STAGE_CALLABLE_BIT_NV`

  - `VK_SHADER_STAGE_CLOSEST_HIT_BIT_NV`

  - `VK_SHADER_STAGE_INTERSECTION_BIT_NV`

  - `VK_SHADER_STAGE_MISS_BIT_NV`

  - `VK_SHADER_STAGE_RAYGEN_BIT_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

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

If [VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html) is supported:

- Extending
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html):

  - `VK_DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-launchid"
  target="_blank" rel="noopener"><code>LaunchIdNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-launchsize"
  target="_blank" rel="noopener"><code>LaunchSizeNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-worldrayorigin"
  target="_blank" rel="noopener"><code>WorldRayOriginNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-worldraydirection"
  target="_blank" rel="noopener"><code>WorldRayDirectionNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-objectrayorigin"
  target="_blank" rel="noopener"><code>ObjectRayOriginNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-objectraydirection"
  target="_blank" rel="noopener"><code>ObjectRayDirectionNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-raytmin"
  target="_blank" rel="noopener"><code>RayTminNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-raytmax"
  target="_blank" rel="noopener"><code>RayTmaxNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-instancecustomindex"
  target="_blank" rel="noopener"><code>InstanceCustomIndexNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-instanceid"
  target="_blank" rel="noopener"><code>InstanceId</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-objecttoworld"
  target="_blank" rel="noopener"><code>ObjectToWorldNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-worldtoobject"
  target="_blank" rel="noopener"><code>WorldToObjectNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-hitt"
  target="_blank" rel="noopener"><code>HitTNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-hitkind"
  target="_blank" rel="noopener"><code>HitKindNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-incomingrayflags"
  target="_blank" rel="noopener"><code>IncomingRayFlagsNV</code></a>

- (modified)`PrimitiveId`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayTracingNV"
  target="_blank" rel="noopener"><code>RayTracingNV</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1\) Are there issues?

**RESOLVED**: Yes.

## <a href="#_sample_code" class="anchor"></a>Sample Code

Example ray generation GLSL shader

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-09-11 (Robert Stepinski, Nuno Subtil, Eric Werness)

  - Internal revisions

- Revision 2, 2018-10-19 (Eric Werness)

  - rename to VK_NV_ray_tracing, add support for callables.

  - too many updates to list

- Revision 3, 2018-11-20 (Daniel Koch)

  - update to use InstanceId instead of InstanceIndex as implemented.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_ray_tracing"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
