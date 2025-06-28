# VkDescriptorImageInfo(3) Manual Page

## Name

VkDescriptorImageInfo - Structure specifying descriptor image information



## [](#_c_specification)C Specification

The `VkDescriptorImageInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorImageInfo {
    VkSampler        sampler;
    VkImageView      imageView;
    VkImageLayout    imageLayout;
} VkDescriptorImageInfo;
```

## [](#_members)Members

- `sampler` is a sampler handle, and is used in descriptor updates for types `VK_DESCRIPTOR_TYPE_SAMPLER` and `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` if the binding being updated does not use immutable samplers.
- `imageView` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or an image view handle, and is used in descriptor updates for types `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`.
- `imageLayout` is the layout that the image subresources accessible from `imageView` will be in at the time this descriptor is accessed. `imageLayout` is used in descriptor updates for types `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`.

## [](#_description)Description

Members of `VkDescriptorImageInfo` that are not used in an update (as described above) are ignored.

Valid Usage

- [](#VUID-VkDescriptorImageInfo-imageView-06712)VUID-VkDescriptorImageInfo-imageView-06712  
  `imageView` **must** not be a 2D array image view created from a 3D image
- [](#VUID-VkDescriptorImageInfo-imageView-07795)VUID-VkDescriptorImageInfo-imageView-07795  
  If `imageView` is a 2D view created from a 3D image, then `descriptorType` **must** be `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`
- [](#VUID-VkDescriptorImageInfo-imageView-07796)VUID-VkDescriptorImageInfo-imageView-07796  
  If `imageView` is a 2D view created from a 3D image, then the image **must** have been created with `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set
- [](#VUID-VkDescriptorImageInfo-descriptorType-06713)VUID-VkDescriptorImageInfo-descriptorType-06713  
  If the [`image2DViewOf3D`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-image2DViewOf3D) feature is not enabled or `descriptorType` is not `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` then `imageView` **must** not be a 2D view created from a 3D image
- [](#VUID-VkDescriptorImageInfo-descriptorType-06714)VUID-VkDescriptorImageInfo-descriptorType-06714  
  If the [`sampler2DViewOf3D`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sampler2DViewOf3D) feature is not enabled or `descriptorType` is not `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` then `imageView` **must** not be a 2D view created from a 3D image
- [](#VUID-VkDescriptorImageInfo-imageView-01976)VUID-VkDescriptorImageInfo-imageView-01976  
  If `imageView` is created from a depth/stencil image, the `aspectMask` used to create the `imageView` **must** include either `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT` but not both
- [](#VUID-VkDescriptorImageInfo-imageLayout-09425)VUID-VkDescriptorImageInfo-imageLayout-09425  
  If `imageLayout` is `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`, then the `aspectMask` used to create `imageView` **must** not include either `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkDescriptorImageInfo-imageLayout-09426)VUID-VkDescriptorImageInfo-imageLayout-09426  
  If `imageLayout` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, then the `aspectMask` used to create `imageView` **must** not include `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkDescriptorImageInfo-imageLayout-00344)VUID-VkDescriptorImageInfo-imageLayout-00344  
  `imageLayout` **must** match the actual [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) of each subresource accessible from `imageView` at the time this descriptor is accessed as defined by the [image layout matching rules](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-layouts-matching-rule)
- [](#VUID-VkDescriptorImageInfo-sampler-01564)VUID-VkDescriptorImageInfo-sampler-01564  
  If `sampler` is used and the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the image is a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), the image **must** have been created with `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`, and the `aspectMask` of the `imageView` **must** be a valid [multi-planar aspect mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar-image-aspect) bit
- [](#VUID-VkDescriptorImageInfo-mutableComparisonSamplers-04450)VUID-VkDescriptorImageInfo-mutableComparisonSamplers-04450  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`mutableComparisonSamplers` is `VK_FALSE`, then `sampler` **must** have been created with [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`compareEnable` set to `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorImageInfo-commonparent)VUID-VkDescriptorImageInfo-commonparent  
  Both of `imageView`, and `sampler` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorDataEXT.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorImageInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0