# VkIndexType(3) Manual Page

## Name

VkIndexType - Type of index buffer indices



## [](#_c_specification)C Specification

Possible values of [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html)::`indexType` and [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer.html)::`indexType`, specifying the size of indices, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkIndexType {
    VK_INDEX_TYPE_UINT16 = 0,
    VK_INDEX_TYPE_UINT32 = 1,
  // Provided by VK_VERSION_1_4
    VK_INDEX_TYPE_UINT8 = 1000265000,
  // Provided by VK_KHR_acceleration_structure
    VK_INDEX_TYPE_NONE_KHR = 1000165000,
  // Provided by VK_NV_ray_tracing
    VK_INDEX_TYPE_NONE_NV = VK_INDEX_TYPE_NONE_KHR,
  // Provided by VK_EXT_index_type_uint8
    VK_INDEX_TYPE_UINT8_EXT = VK_INDEX_TYPE_UINT8,
  // Provided by VK_KHR_index_type_uint8
    VK_INDEX_TYPE_UINT8_KHR = VK_INDEX_TYPE_UINT8,
} VkIndexType;
```

## [](#_description)Description

- `VK_INDEX_TYPE_UINT16` specifies that indices are 16-bit unsigned integer values.
- `VK_INDEX_TYPE_UINT32` specifies that indices are 32-bit unsigned integer values.
- `VK_INDEX_TYPE_NONE_KHR` specifies that no indices are provided.
- `VK_INDEX_TYPE_UINT8` specifies that indices are 8-bit unsigned integer values.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html), [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html), [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html), [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html), [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html), [VkBindIndexBufferIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindIndexBufferIndirectCommandEXT.html), [VkBindIndexBufferIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindIndexBufferIndirectCommandNV.html), [VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTrianglesNV.html), [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenNV.html), [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer.html), [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html), [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndexType)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0