# VkVertexInputAttributeDescription(3) Manual Page

## Name

VkVertexInputAttributeDescription - Structure specifying vertex input
attribute description



## <a href="#_c_specification" class="anchor"></a>C Specification

Each vertex input attribute is specified by the
`VkVertexInputAttributeDescription` structure, defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkVertexInputAttributeDescription {
    uint32_t    location;
    uint32_t    binding;
    VkFormat    format;
    uint32_t    offset;
} VkVertexInputAttributeDescription;
```

## <a href="#_members" class="anchor"></a>Members

- `location` is the shader input location number for this attribute.

- `binding` is the binding number which this attribute takes its data
  from.

- `format` is the size and type of the vertex attribute data.

- `offset` is a byte offset of this attribute relative to the start of
  an element in the vertex input binding.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVertexInputAttributeDescription-location-00620"
  id="VUID-VkVertexInputAttributeDescription-location-00620"></a>
  VUID-VkVertexInputAttributeDescription-location-00620  
  `location` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputAttributes`

- <a href="#VUID-VkVertexInputAttributeDescription-binding-00621"
  id="VUID-VkVertexInputAttributeDescription-binding-00621"></a>
  VUID-VkVertexInputAttributeDescription-binding-00621  
  `binding` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-VkVertexInputAttributeDescription-offset-00622"
  id="VUID-VkVertexInputAttributeDescription-offset-00622"></a>
  VUID-VkVertexInputAttributeDescription-offset-00622  
  `offset` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxVertexInputAttributeOffset`

- <a href="#VUID-VkVertexInputAttributeDescription-format-00623"
  id="VUID-VkVertexInputAttributeDescription-format-00623"></a>
  VUID-VkVertexInputAttributeDescription-format-00623  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-view-format-features"
  target="_blank" rel="noopener">format features</a> of `format`
  **must** contain `VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT`

- <a
  href="#VUID-VkVertexInputAttributeDescription-vertexAttributeAccessBeyondStride-04457"
  id="VUID-VkVertexInputAttributeDescription-vertexAttributeAccessBeyondStride-04457"></a>
  VUID-VkVertexInputAttributeDescription-vertexAttributeAccessBeyondStride-04457  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`vertexAttributeAccessBeyondStride`
  is `VK_FALSE`, the sum of `offset` plus the size of the vertex
  attribute data described by `format` **must** not be greater than
  `stride` in the
  [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription.html)
  referenced in `binding`

Valid Usage (Implicit)

- <a href="#VUID-VkVertexInputAttributeDescription-format-parameter"
  id="VUID-VkVertexInputAttributeDescription-format-parameter"></a>
  VUID-VkVertexInputAttributeDescription-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVertexInputAttributeDescription"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
