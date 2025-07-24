# vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(3) Manual Page

## Name

vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI - Query maximum supported subpass shading workgroup size for a give render pass



## [](#_c_specification)C Specification

A subpass shading pipeline’s workgroup size is a 2D vector with number of power-of-two in width and height. The maximum number of width and height is implementation-dependent, and **may** vary for different formats and sample counts of attachments in a render pass.

To query the maximum workgroup size, call:

```c++
// Provided by VK_HUAWEI_subpass_shading
VkResult vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(
    VkDevice                                    device,
    VkRenderPass                                renderpass,
    VkExtent2D*                                 pMaxWorkgroupSize);
```

## [](#_parameters)Parameters

- `device` is a handle to a local device object that was used to create the given render pass.
- `renderpass` is a handle to a render pass object describing the environment in which the pipeline will be used. The pipeline **must** only be used with a render pass instance compatible with the one provided. See [Render Pass Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-compatibility) for more information.
- `pMaxWorkgroupSize` is a pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-device-parameter)VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parameter)VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parameter  
  `renderpass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) handle
- [](#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-pMaxWorkgroupSize-parameter)VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-pMaxWorkgroupSize-parameter  
  `pMaxWorkgroupSize` **must** be a valid pointer to [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) structures
- [](#VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parent)VUID-vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI-renderpass-parent  
  `renderpass` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_HUAWEI\_subpass\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_subpass_shading.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0