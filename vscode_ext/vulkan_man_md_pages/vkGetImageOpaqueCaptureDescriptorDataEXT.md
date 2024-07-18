# vkGetImageOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetImageOpaqueCaptureDescriptorDataEXT - Get image opaque capture
descriptor data



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the opaque capture descriptor data for an image, call:

``` c
// Provided by VK_EXT_descriptor_buffer
VkResult vkGetImageOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkImageCaptureDescriptorDataInfoEXT*  pInfo,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that gets the data.

- `pInfo` is a pointer to a
  [VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCaptureDescriptorDataInfoEXT.html)
  structure specifying the image.

- `pData` is a pointer to an application-allocated buffer where the data
  will be written.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-None-08076"
  id="VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-None-08076"></a>
  VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-None-08076  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-08077"
  id="VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-08077"></a>
  VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-08077  
  `pData` **must** point to a buffer that is at least
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`imageCaptureReplayDescriptorDataSize`
  bytes in size

- <a href="#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-08078"
  id="VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-08078"></a>
  VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-08078  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-parameter"
  id="VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-parameter"></a>
  VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pInfo-parameter"
  id="VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pInfo-parameter"></a>
  VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCaptureDescriptorDataInfoEXT.html)
  structure

- <a href="#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-parameter"
  id="VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-parameter"></a>
  VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-parameter  
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
[VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCaptureDescriptorDataInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageOpaqueCaptureDescriptorDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
