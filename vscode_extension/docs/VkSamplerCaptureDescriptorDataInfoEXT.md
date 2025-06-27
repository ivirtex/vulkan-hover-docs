# VkSamplerCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkSamplerCaptureDescriptorDataInfoEXT - Structure specifying a sampler
for descriptor capture



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the sampler to get descriptor buffer capture data for
is passed in a `VkSamplerCaptureDescriptorDataInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkSamplerCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkSampler          sampler;
} VkSamplerCaptureDescriptorDataInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `sampler` is the `VkSampler` handle of the sampler to get opaque
  capture data for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-08087"
  id="VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-08087"></a>
  VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-08087  
  `sampler` **must** have been created with
  `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`flags`

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerCaptureDescriptorDataInfoEXT-sType-sType"
  id="VUID-VkSamplerCaptureDescriptorDataInfoEXT-sType-sType"></a>
  VUID-VkSamplerCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`

- <a href="#VUID-VkSamplerCaptureDescriptorDataInfoEXT-pNext-pNext"
  id="VUID-VkSamplerCaptureDescriptorDataInfoEXT-pNext-pNext"></a>
  VUID-VkSamplerCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-parameter"
  id="VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-parameter"></a>
  VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-parameter  
  `sampler` **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetSamplerOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSamplerOpaqueCaptureDescriptorDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerCaptureDescriptorDataInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
