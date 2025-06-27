# vkGetBufferOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetBufferOpaqueCaptureDescriptorDataEXT - Get buffer opaque capture
descriptor data



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the opaque descriptor data for a buffer, call:

``` c
// Provided by VK_EXT_descriptor_buffer
VkResult vkGetBufferOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkBufferCaptureDescriptorDataInfoEXT* pInfo,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that gets the data.

- `pInfo` is a pointer to a
  [VkBufferCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCaptureDescriptorDataInfoEXT.html)
  structure specifying the buffer.

- `pData` is a pointer to an application-allocated buffer where the data
  will be written.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-None-08072"
  id="VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-None-08072"></a>
  VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-None-08072  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-08073"
  id="VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-08073"></a>
  VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-08073  
  `pData` **must** point to a buffer that is at least
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`bufferCaptureReplayDescriptorDataSize`
  bytes in size

- <a href="#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-08074"
  id="VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-08074"></a>
  VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-08074  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-parameter"
  id="VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-parameter"></a>
  VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pInfo-parameter"
  id="VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pInfo-parameter"></a>
  VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkBufferCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCaptureDescriptorDataInfoEXT.html)
  structure

- <a
  href="#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-parameter"
  id="VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-parameter"></a>
  VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-parameter  
  `pData` **must** be a pointer value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkBufferCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCaptureDescriptorDataInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetBufferOpaqueCaptureDescriptorDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
