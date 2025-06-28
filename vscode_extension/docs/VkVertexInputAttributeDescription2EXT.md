# VkVertexInputAttributeDescription2EXT(3) Manual Page

## Name

VkVertexInputAttributeDescription2EXT - Structure specifying the extended vertex input attribute description



## [](#_c_specification)C Specification

The `VkVertexInputAttributeDescription2EXT` structure is defined as:

```c++
// Provided by VK_EXT_shader_object, VK_EXT_vertex_input_dynamic_state
typedef struct VkVertexInputAttributeDescription2EXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           location;
    uint32_t           binding;
    VkFormat           format;
    uint32_t           offset;
} VkVertexInputAttributeDescription2EXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `location` is the shader input location number for this attribute.
- `binding` is the binding number which this attribute takes its data from.
- `format` is the size and type of the vertex attribute data.
- `offset` is a byte offset of this attribute relative to the start of an element in the vertex input binding.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVertexInputAttributeDescription2EXT-location-06228)VUID-VkVertexInputAttributeDescription2EXT-location-06228  
  `location` **must** be less than `VkPhysicalDeviceLimits`::`maxVertexInputAttributes`
- [](#VUID-VkVertexInputAttributeDescription2EXT-binding-06229)VUID-VkVertexInputAttributeDescription2EXT-binding-06229  
  `binding` **must** be less than `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-VkVertexInputAttributeDescription2EXT-offset-06230)VUID-VkVertexInputAttributeDescription2EXT-offset-06230  
  `offset` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputAttributeOffset`
- [](#VUID-VkVertexInputAttributeDescription2EXT-format-04805)VUID-VkVertexInputAttributeDescription2EXT-format-04805  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `format` **must** contain `VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT`
- [](#VUID-VkVertexInputAttributeDescription2EXT-vertexAttributeAccessBeyondStride-04806)VUID-VkVertexInputAttributeDescription2EXT-vertexAttributeAccessBeyondStride-04806  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`vertexAttributeAccessBeyondStride` is `VK_FALSE`, the sum of `offset` plus the size of the vertex attribute data described by `format` **must** not be greater than `stride` in the [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html) referenced in `binding`

Valid Usage (Implicit)

- [](#VUID-VkVertexInputAttributeDescription2EXT-sType-sType)VUID-VkVertexInputAttributeDescription2EXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VERTEX_INPUT_ATTRIBUTE_DESCRIPTION_2_EXT`
- [](#VUID-VkVertexInputAttributeDescription2EXT-format-parameter)VUID-VkVertexInputAttributeDescription2EXT-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_EXT\_vertex\_input\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_input_dynamic_state.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVertexInputAttributeDescription2EXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0