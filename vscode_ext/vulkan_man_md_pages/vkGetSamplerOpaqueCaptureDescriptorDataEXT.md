# vkGetSamplerOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetSamplerOpaqueCaptureDescriptorDataEXT - Get sampler opaque capture
descriptor data



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the opaque capture descriptor data for a sampler, call:

``` c
// Provided by VK_EXT_descriptor_buffer
VkResult vkGetSamplerOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkSamplerCaptureDescriptorDataInfoEXT* pInfo,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that gets the data.

- `pInfo` is a pointer to a
  [VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html)
  structure specifying the sampler.

- `pData` is a pointer to a user-allocated buffer where the data will be
  written.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-None-08084"
  id="VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-None-08084"></a>
  VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-None-08084  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-08085"
  id="VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-08085"></a>
  VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-08085  
  `pData` **must** point to a buffer that is at least
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`samplerCaptureReplayDescriptorDataSize`
  bytes in size

- <a href="#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-08086"
  id="VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-08086"></a>
  VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-08086  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-parameter"
  id="VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-parameter"></a>
  VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pInfo-parameter"
  id="VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pInfo-parameter"></a>
  VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html)
  structure

- <a
  href="#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-parameter"
  id="VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-parameter"></a>
  VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-parameter  
  `pData` **must** be a pointer value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetSamplerOpaqueCaptureDescriptorDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
