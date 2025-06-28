# VkVertexInputAttributeDescription(3) Manual Page

## Name

VkVertexInputAttributeDescription - Structure specifying vertex input attribute description



## [](#_c_specification)C Specification

Each vertex input attribute is specified by the `VkVertexInputAttributeDescription` structure, defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkVertexInputAttributeDescription {
    uint32_t    location;
    uint32_t    binding;
    VkFormat    format;
    uint32_t    offset;
} VkVertexInputAttributeDescription;
```

## [](#_members)Members

- `location` is the shader input location number for this attribute.
- `binding` is the binding number which this attribute takes its data from.
- `format` is the size and type of the vertex attribute data.
- `offset` is a byte offset of this attribute relative to the start of an element in the vertex input binding.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVertexInputAttributeDescription-location-00620)VUID-VkVertexInputAttributeDescription-location-00620  
  `location` **must** be less than `VkPhysicalDeviceLimits`::`maxVertexInputAttributes`
- [](#VUID-VkVertexInputAttributeDescription-binding-00621)VUID-VkVertexInputAttributeDescription-binding-00621  
  `binding` **must** be less than `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-VkVertexInputAttributeDescription-offset-00622)VUID-VkVertexInputAttributeDescription-offset-00622  
  `offset` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputAttributeOffset`
- [](#VUID-VkVertexInputAttributeDescription-format-00623)VUID-VkVertexInputAttributeDescription-format-00623  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `format` **must** contain `VK_FORMAT_FEATURE_VERTEX_BUFFER_BIT`
- [](#VUID-VkVertexInputAttributeDescription-vertexAttributeAccessBeyondStride-04457)VUID-VkVertexInputAttributeDescription-vertexAttributeAccessBeyondStride-04457  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`vertexAttributeAccessBeyondStride` is `VK_FALSE`, the sum of `offset` plus the size of the vertex attribute data described by `format` **must** not be greater than `stride` in the [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html) referenced in `binding`

Valid Usage (Implicit)

- [](#VUID-VkVertexInputAttributeDescription-format-parameter)VUID-VkVertexInputAttributeDescription-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVertexInputAttributeDescription)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0