# VkDescriptorImageInfo(3) Manual Page

## Name

VkDescriptorImageInfo - Structure specifying descriptor image
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDescriptorImageInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorImageInfo {
    VkSampler        sampler;
    VkImageView      imageView;
    VkImageLayout    imageLayout;
} VkDescriptorImageInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sampler` is a sampler handle, and is used in descriptor updates for
  types `VK_DESCRIPTOR_TYPE_SAMPLER` and
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` if the binding being
  updated does not use immutable samplers.

- `imageView` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or an image view
  handle, and is used in descriptor updates for types
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`.

- `imageLayout` is the layout that the image subresources accessible
  from `imageView` will be in at the time this descriptor is accessed.
  `imageLayout` is used in descriptor updates for types
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`.

## <a href="#_description" class="anchor"></a>Description

Members of `VkDescriptorImageInfo` that are not used in an update (as
described above) are ignored.

Valid Usage

- <a href="#VUID-VkDescriptorImageInfo-imageView-06712"
  id="VUID-VkDescriptorImageInfo-imageView-06712"></a>
  VUID-VkDescriptorImageInfo-imageView-06712  
  `imageView` **must** not be a 2D array image view created from a 3D
  image

- <a href="#VUID-VkDescriptorImageInfo-imageView-07795"
  id="VUID-VkDescriptorImageInfo-imageView-07795"></a>
  VUID-VkDescriptorImageInfo-imageView-07795  
  If `imageView` is a 2D view created from a 3D image, then
  `descriptorType` **must** be `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`

- <a href="#VUID-VkDescriptorImageInfo-imageView-07796"
  id="VUID-VkDescriptorImageInfo-imageView-07796"></a>
  VUID-VkDescriptorImageInfo-imageView-07796  
  If `imageView` is a 2D view created from a 3D image, then the image
  **must** have been created with
  `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set

- <a href="#VUID-VkDescriptorImageInfo-descriptorType-06713"
  id="VUID-VkDescriptorImageInfo-descriptorType-06713"></a>
  VUID-VkDescriptorImageInfo-descriptorType-06713  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-image2DViewOf3D"
  target="_blank" rel="noopener"><code>image2DViewOf3D</code></a>
  feature is not enabled or `descriptorType` is not
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` then `imageView` **must** not be a
  2D view created from a 3D image

- <a href="#VUID-VkDescriptorImageInfo-descriptorType-06714"
  id="VUID-VkDescriptorImageInfo-descriptorType-06714"></a>
  VUID-VkDescriptorImageInfo-descriptorType-06714  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sampler2DViewOf3D"
  target="_blank" rel="noopener"><code>sampler2DViewOf3D</code></a>
  feature is not enabled or `descriptorType` is not
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` then `imageView` **must**
  not be a 2D view created from a 3D image

- <a href="#VUID-VkDescriptorImageInfo-imageView-01976"
  id="VUID-VkDescriptorImageInfo-imageView-01976"></a>
  VUID-VkDescriptorImageInfo-imageView-01976  
  If `imageView` is created from a depth/stencil image, the `aspectMask`
  used to create the `imageView` **must** include either
  `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT` but not
  both

- <a href="#VUID-VkDescriptorImageInfo-imageLayout-09425"
  id="VUID-VkDescriptorImageInfo-imageLayout-09425"></a>
  VUID-VkDescriptorImageInfo-imageLayout-09425  
  If `imageLayout` is `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`, then
  the `aspectMask` used to create `imageView` **must** not include
  either `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkDescriptorImageInfo-imageLayout-09426"
  id="VUID-VkDescriptorImageInfo-imageLayout-09426"></a>
  VUID-VkDescriptorImageInfo-imageLayout-09426  
  If `imageLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, then the
  `aspectMask` used to create `imageView` **must** not include
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkDescriptorImageInfo-imageLayout-00344"
  id="VUID-VkDescriptorImageInfo-imageLayout-00344"></a>
  VUID-VkDescriptorImageInfo-imageLayout-00344  
  `imageLayout` **must** match the actual
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) of each subresource accessible
  from `imageView` at the time this descriptor is accessed as defined by
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-layouts-matching-rule"
  target="_blank" rel="noopener">image layout matching rules</a>

- <a href="#VUID-VkDescriptorImageInfo-sampler-01564"
  id="VUID-VkDescriptorImageInfo-sampler-01564"></a>
  VUID-VkDescriptorImageInfo-sampler-01564  
  If `sampler` is used and the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of the image is
  a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar format</a>, the image
  **must** have been created with `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`,
  and the `aspectMask` of the `imageView` **must** be a valid <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-planes-image-aspect"
  target="_blank" rel="noopener">multi-planar aspect mask</a> bit

- <a href="#VUID-VkDescriptorImageInfo-mutableComparisonSamplers-04450"
  id="VUID-VkDescriptorImageInfo-mutableComparisonSamplers-04450"></a>
  VUID-VkDescriptorImageInfo-mutableComparisonSamplers-04450  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`mutableComparisonSamplers`
  is `VK_FALSE`, then `sampler` **must** have been created with
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`compareEnable` set
  to `VK_FALSE`

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorImageInfo-commonparent"
  id="VUID-VkDescriptorImageInfo-commonparent"></a>
  VUID-VkDescriptorImageInfo-commonparent  
  Both of `imageView`, and `sampler` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorDataEXT.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html),
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorImageInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
