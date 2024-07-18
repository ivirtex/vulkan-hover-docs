# VkVertexInputAttributeDescription2EXT(3) Manual Page

## Name

VkVertexInputAttributeDescription2EXT - Structure specifying the
extended vertex input attribute description



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVertexInputAttributeDescription2EXT` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `location` is the shader input location number for this attribute.

- `binding` is the binding number which this attribute takes its data
  from.

- `format` is the size and type of the vertex attribute data.

- `offset` is a byte offset of this attribute relative to the start of
  an element in the vertex input binding.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVertexInputAttributeDescription2EXT-location-06228"
  id="VUID-VkVertexInputAttributeDescription2EXT-location-06228"></a>
  VUID-VkVertexInputAttributeDescription2EXT-location-06228  
  `location` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputAttributes`

- <a href="#VUID-VkVertexInputAttributeDescription2EXT-binding-06229"
  id="VUID-VkVertexInputAttributeDescription2EXT-binding-06229"></a>
  VUID-VkVertexInputAttributeDescription2EXT-binding-06229  
  `binding` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-VkVertexInputAttributeDescription2EXT-offset-06230"
  id="VUID-VkVertexInputAttributeDescription2EXT-offset-06230"></a>
  VUID-VkVertexInputAttributeDescription2EXT-offset-06230  
  `offset` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxVertexInputAttributeOffset`

- <a href="#VUID-VkVertexInputAttributeDescription2EXT-format-04805"
  id="VUID-VkVertexInputAttributeDescription2EXT-format-04805"></a>
  VUID-VkVertexInputAttributeDescription2EXT-format-04805  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-view-format-features"
  target="_blank" rel="noopener">format features</a> of `format`
  **must** contain `VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT`

- <a
  href="#VUID-VkVertexInputAttributeDescription2EXT-vertexAttributeAccessBeyondStride-04806"
  id="VUID-VkVertexInputAttributeDescription2EXT-vertexAttributeAccessBeyondStride-04806"></a>
  VUID-VkVertexInputAttributeDescription2EXT-vertexAttributeAccessBeyondStride-04806  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`vertexAttributeAccessBeyondStride`
  is `VK_FALSE`, the sum of `offset` plus the size of the vertex
  attribute data described by `format` **must** not be greater than
  `stride` in the
  [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)
  referenced in `binding`

Valid Usage (Implicit)

- <a href="#VUID-VkVertexInputAttributeDescription2EXT-sType-sType"
  id="VUID-VkVertexInputAttributeDescription2EXT-sType-sType"></a>
  VUID-VkVertexInputAttributeDescription2EXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VERTEX_INPUT_ATTRIBUTE_DESCRIPTION_2_EXT`

- <a href="#VUID-VkVertexInputAttributeDescription2EXT-format-parameter"
  id="VUID-VkVertexInputAttributeDescription2EXT-format-parameter"></a>
  VUID-VkVertexInputAttributeDescription2EXT-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VK_EXT_vertex_input_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_input_dynamic_state.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVertexInputAttributeDescription2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
