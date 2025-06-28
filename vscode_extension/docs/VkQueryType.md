# VkQueryType(3) Manual Page

## Name

VkQueryType - Specify the type of queries managed by a query pool



## [](#_c_specification)C Specification

Possible values of [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html)::`queryType`, specifying the type of queries managed by the pool, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkQueryType {
    VK_QUERY_TYPE_OCCLUSION = 0,
    VK_QUERY_TYPE_PIPELINE_STATISTICS = 1,
    VK_QUERY_TYPE_TIMESTAMP = 2,
  // Provided by VK_KHR_video_queue
    VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR = 1000023000,
  // Provided by VK_EXT_transform_feedback
    VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT = 1000028004,
  // Provided by VK_KHR_performance_query
    VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR = 1000116000,
  // Provided by VK_KHR_acceleration_structure
    VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR = 1000150000,
  // Provided by VK_KHR_acceleration_structure
    VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR = 1000150001,
  // Provided by VK_NV_ray_tracing
    VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV = 1000165000,
  // Provided by VK_INTEL_performance_query
    VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL = 1000210000,
  // Provided by VK_KHR_video_encode_queue
    VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR = 1000299000,
  // Provided by VK_EXT_mesh_shader
    VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT = 1000328000,
  // Provided by VK_EXT_primitives_generated_query
    VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT = 1000382000,
  // Provided by VK_KHR_ray_tracing_maintenance1
    VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR = 1000386000,
  // Provided by VK_KHR_ray_tracing_maintenance1
    VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR = 1000386001,
  // Provided by VK_EXT_opacity_micromap
    VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT = 1000396000,
  // Provided by VK_EXT_opacity_micromap
    VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT = 1000396001,
} VkQueryType;
```

## [](#_description)Description

- `VK_QUERY_TYPE_OCCLUSION` specifies an [occlusion query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-occlusion).
- `VK_QUERY_TYPE_PIPELINE_STATISTICS` specifies a [pipeline statistics query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-pipestats).
- `VK_QUERY_TYPE_TIMESTAMP` specifies a [timestamp query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-timestamps).
- `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` specifies a [performance query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-performance).
- `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` specifies a [transform feedback query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-transform-feedback).
- `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` specifies a [primitives generated query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-primitives-generated).
- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR` specifies a [acceleration structure size query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-copying) for use with [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html) or [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteAccelerationStructuresPropertiesKHR.html).
- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR` specifies a [serialization acceleration structure size query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-copying).
- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` specifies an [acceleration structure size query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-copying) for use with [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html) or [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteAccelerationStructuresPropertiesKHR.html).
- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR` specifies a [serialization acceleration structure pointer count query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#serialized-as-header).
- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV` specifies an [acceleration structure size query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-copying) for use with [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html).
- `VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL` specifies a [Intel performance query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-performance-intel).
- `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` specifies a [result status query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-result-status-only).
- `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR` specifies a [video encode feedback query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-video-encode-feedback).
- `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT` specifies a [generated mesh primitives query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-mesh-shader).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html), [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html), [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html), [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteMicromapsPropertiesEXT.html), [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteAccelerationStructuresPropertiesKHR.html), [vkWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteMicromapsPropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryType)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0