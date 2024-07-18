# VkQueryType(3) Manual Page

## Name

VkQueryType - Specify the type of queries managed by a query pool



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html)::`queryType`,
specifying the type of queries managed by the pool, are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_QUERY_TYPE_OCCLUSION` specifies an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-occlusion"
  target="_blank" rel="noopener">occlusion query</a>.

- `VK_QUERY_TYPE_PIPELINE_STATISTICS` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-pipestats"
  target="_blank" rel="noopener">pipeline statistics query</a>.

- `VK_QUERY_TYPE_TIMESTAMP` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-timestamps"
  target="_blank" rel="noopener">timestamp query</a>.

- `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-performance"
  target="_blank" rel="noopener">performance query</a>.

- `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-transform-feedback"
  target="_blank" rel="noopener">transform feedback query</a>.

- `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-primitives-generated"
  target="_blank" rel="noopener">primitives generated query</a>.

- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR` specifies a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-copying"
  target="_blank" rel="noopener">acceleration structure size query</a>
  for use with
  [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html)
  or
  [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteAccelerationStructuresPropertiesKHR.html).

- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`
  specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-copying"
  target="_blank" rel="noopener">serialization acceleration structure size
  query</a>.

- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` specifies an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-copying"
  target="_blank" rel="noopener">acceleration structure size query</a>
  for use with
  [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html)
  or
  [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteAccelerationStructuresPropertiesKHR.html).

- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`
  specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#serialized-as-header"
  target="_blank" rel="noopener">serialization acceleration structure
  pointer count query</a>.

- `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV` specifies an
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-copying"
  target="_blank" rel="noopener">acceleration structure size query</a>
  for use with
  [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html).

- `VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-performance-intel"
  target="_blank" rel="noopener">Intel performance query</a>.

- `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-result-status-only"
  target="_blank" rel="noopener">result status query</a>.

- `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-video-encode-feedback"
  target="_blank" rel="noopener">video encode feedback query</a>.

- `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-mesh-shader"
  target="_blank" rel="noopener">generated mesh primitives query</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html),
[vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html),
[vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html),
[vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteMicromapsPropertiesEXT.html),
[vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteAccelerationStructuresPropertiesKHR.html),
[vkWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteMicromapsPropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
