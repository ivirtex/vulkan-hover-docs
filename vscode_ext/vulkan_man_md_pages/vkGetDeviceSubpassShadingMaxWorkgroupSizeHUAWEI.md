# vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(3) Manual Page

## Name

vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI - Query maximum
supported subpass shading workgroup size for a give render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

A subpass shading pipeline’s workgroup size is a 2D vector with number
of power-of-two in width and height. The maximum number of width and
height is implementation-dependent, and **may** vary for different
formats and sample counts of attachments in a render pass.

To query the maximum workgroup size, call:

``` c
// Provided by VK_HUAWEI_subpass_shading
VkResult vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(
    VkDevice                                    device,
    VkRenderPass                                renderpass,
    VkExtent2D*                                 pMaxWorkgroupSize);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is a handle to a local device object that was used to create
  the given render pass.

- `renderPass` is a handle to a render pass object describing the
  environment in which the pipeline will be used. The pipeline **must**
  only be used with a render pass instance compatible with the one
  provided. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-compatibility"
  target="_blank" rel="noopener">Render Pass Compatibility</a> for more
  information.

- `pMaxWorkgroupSize` is a pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html)
  structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-device-parameter"
  id="VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-device-parameter"></a>
  VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parameter"
  id="VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parameter"></a>
  VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parameter  
  `renderpass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

- <a
  href="#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-pMaxWorkgroupSize-parameter"
  id="VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-pMaxWorkgroupSize-parameter"></a>
  VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-pMaxWorkgroupSize-parameter  
  `pMaxWorkgroupSize` **must** be a valid pointer to
  [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) structures

- <a
  href="#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parent"
  id="VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parent"></a>
  VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parent  
  `renderpass` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_subpass_shading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_subpass_shading.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
