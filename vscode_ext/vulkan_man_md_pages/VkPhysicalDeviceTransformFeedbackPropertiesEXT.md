# VkPhysicalDeviceTransformFeedbackPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceTransformFeedbackPropertiesEXT - Structure describing
transform feedback properties that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceTransformFeedbackPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_transform_feedback
typedef struct VkPhysicalDeviceTransformFeedbackPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxTransformFeedbackStreams;
    uint32_t           maxTransformFeedbackBuffers;
    VkDeviceSize       maxTransformFeedbackBufferSize;
    uint32_t           maxTransformFeedbackStreamDataSize;
    uint32_t           maxTransformFeedbackBufferDataSize;
    uint32_t           maxTransformFeedbackBufferDataStride;
    VkBool32           transformFeedbackQueries;
    VkBool32           transformFeedbackStreamsLinesTriangles;
    VkBool32           transformFeedbackRasterizationStreamSelect;
    VkBool32           transformFeedbackDraw;
} VkPhysicalDeviceTransformFeedbackPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxTransformFeedbackStreams"></span>
  `maxTransformFeedbackStreams` is the maximum number of vertex streams
  that can be output from geometry shaders declared with the
  `GeometryStreams` capability. If the implementation does not support
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`geometryStreams` then
  `maxTransformFeedbackStreams` **must** be set to `1`.

- <span id="limits-maxTransformFeedbackBuffers"></span>
  `maxTransformFeedbackBuffers` is the maximum number of transform
  feedback buffers that can be bound for capturing shader outputs from
  the last <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a>.

- <span id="limits-maxTransformFeedbackBufferSize"></span>
  `maxTransformFeedbackBufferSize` is the maximum size that can be
  specified when binding a buffer for transform feedback in
  [vkCmdBindTransformFeedbackBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindTransformFeedbackBuffersEXT.html).

- <span id="limits-maxTransformFeedbackStreamDataSize"></span>
  `maxTransformFeedbackStreamDataSize` is the maximum amount of data in
  bytes for each vertex that captured to one or more transform feedback
  buffers associated with a specific vertex stream.

- <span id="limits-maxTransformFeedbackBufferDataSize"></span>
  `maxTransformFeedbackBufferDataSize` is the maximum amount of data in
  bytes for each vertex that can be captured to a specific transform
  feedback buffer.

- <span id="limits-maxTransformFeedbackBufferDataStride"></span>
  `maxTransformFeedbackBufferDataStride` is the maximum stride between
  each capture of vertex data to the buffer.

- <span id="limits-transformFeedbackQueries"></span>
  `transformFeedbackQueries` is `VK_TRUE` if the implementation supports
  the `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` query type.
  `transformFeedbackQueries` is `VK_FALSE` if queries of this type
  **cannot** be created.

- <span id="limits-transformFeedbackStreamsLinesTriangles"></span>
  `transformFeedbackStreamsLinesTriangles` is `VK_TRUE` if the
  implementation supports the geometry shader `OpExecutionMode` of
  `OutputLineStrip` and `OutputTriangleStrip` in addition to
  `OutputPoints` when more than one vertex stream is output. If
  `transformFeedbackStreamsLinesTriangles` is `VK_FALSE` the
  implementation only supports an `OpExecutionMode` of `OutputPoints`
  when more than one vertex stream is output from the geometry shader.

- <span id="limits-transformFeedbackRasterizationStreamSelect"></span>
  `transformFeedbackRasterizationStreamSelect` is `VK_TRUE` if the
  implementation supports the `GeometryStreams` SPIR-V capability and
  the application can use
  [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)
  to modify which vertex stream output is used for rasterization.
  Otherwise vertex stream `0` **must** always be used for rasterization.

- <span id="limits-transformFeedbackDraw"></span>
  `transformFeedbackDraw` is `VK_TRUE` if the implementation supports
  the
  [vkCmdDrawIndirectByteCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirectByteCountEXT.html)
  function otherwise the function **must** not be called.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceTransformFeedbackPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceTransformFeedbackPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceTransformFeedbackPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceTransformFeedbackPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceTransformFeedbackPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
