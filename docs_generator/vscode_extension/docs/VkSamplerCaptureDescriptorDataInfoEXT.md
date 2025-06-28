# VkSamplerCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkSamplerCaptureDescriptorDataInfoEXT - Structure specifying a sampler for descriptor capture



## [](#_c_specification)C Specification

Information about the sampler to get descriptor buffer capture data for is passed in a `VkSamplerCaptureDescriptorDataInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkSamplerCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkSampler          sampler;
} VkSamplerCaptureDescriptorDataInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `sampler` is the `VkSampler` handle of the sampler to get opaque capture data for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-08087)VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-08087  
  `sampler` **must** have been created with `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`flags`

Valid Usage (Implicit)

- [](#VUID-VkSamplerCaptureDescriptorDataInfoEXT-sType-sType)VUID-VkSamplerCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
- [](#VUID-VkSamplerCaptureDescriptorDataInfoEXT-pNext-pNext)VUID-VkSamplerCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-parameter)VUID-VkSamplerCaptureDescriptorDataInfoEXT-sampler-parameter  
  `sampler` **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) handle

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetSamplerOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSamplerOpaqueCaptureDescriptorDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerCaptureDescriptorDataInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0