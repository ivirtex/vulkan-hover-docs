# VkSampler(3) Manual Page

## Name

VkSampler - Opaque handle to a sampler object



## [](#_c_specification)C Specification

`VkSampler` objects represent the state of an image sampler which is used by the implementation to read image data and apply filtering and other transformations for the shader.

Samplers are represented by `VkSampler` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSampler)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorDataEXT.html), [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html), [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html), [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewHandleInfoNVX.html), [VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html), [vkCreateSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSampler.html), [vkDestroySampler](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySampler.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSampler)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0