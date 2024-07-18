# VkIndexType(3) Manual Page

## Name

VkIndexType - Type of index buffer indices



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html)::`indexType`
and [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html)::`indexType`,
specifying the size of indices, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkIndexType {
    VK_INDEX_TYPE_UINT16 = 0,
    VK_INDEX_TYPE_UINT32 = 1,
  // Provided by VK_KHR_acceleration_structure
    VK_INDEX_TYPE_NONE_KHR = 1000165000,
  // Provided by VK_KHR_index_type_uint8
    VK_INDEX_TYPE_UINT8_KHR = 1000265000,
  // Provided by VK_NV_ray_tracing
    VK_INDEX_TYPE_NONE_NV = VK_INDEX_TYPE_NONE_KHR,
  // Provided by VK_EXT_index_type_uint8
    VK_INDEX_TYPE_UINT8_EXT = VK_INDEX_TYPE_UINT8_KHR,
} VkIndexType;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_INDEX_TYPE_UINT16` specifies that indices are 16-bit unsigned
  integer values.

- `VK_INDEX_TYPE_UINT32` specifies that indices are 32-bit unsigned
  integer values.

- `VK_INDEX_TYPE_NONE_KHR` specifies that no indices are provided.

- `VK_INDEX_TYPE_UINT8_KHR` specifies that indices are 8-bit unsigned
  integer values.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html),
[VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html),
[VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html),
[VkBindIndexBufferIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindIndexBufferIndirectCommandNV.html),
[VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTrianglesNV.html),
[VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutTokenNV.html),
[vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html),
[vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndexType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
