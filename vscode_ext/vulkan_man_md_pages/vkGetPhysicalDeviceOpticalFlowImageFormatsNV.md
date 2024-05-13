# vkGetPhysicalDeviceOpticalFlowImageFormatsNV(3) Manual Page

## Name

vkGetPhysicalDeviceOpticalFlowImageFormatsNV - Query image formats for
optical flow



## <a href="#_c_specification" class="anchor"></a>C Specification

To enumerate the supported image formats for a specific optical flow
usage, call:

``` c
// Provided by VK_NV_optical_flow
VkResult vkGetPhysicalDeviceOpticalFlowImageFormatsNV(
    VkPhysicalDevice                            physicalDevice,
    const VkOpticalFlowImageFormatInfoNV*       pOpticalFlowImageFormatInfo,
    uint32_t*                                   pFormatCount,
    VkOpticalFlowImageFormatPropertiesNV*       pImageFormatProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device being queried.

- <span id="opticalflow-getimageformat-pOpticalFlowImageFormatInfo"></span>
  `pOpticalFlowImageFormatInfo` is a pointer to a
  [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html)
  structure specifying the optical flow usage for which information is
  returned.

- <span id="opticalflow-getimageformat-pFormatCount"></span>
  `pFormatCount` is a pointer to an integer related to the number of
  optical flow properties available or queried, as described below.

- <span id="opticalflow-getimageformat-pImageFormatProperties"></span>
  `pImageFormatProperties` is a pointer to an array of
  [VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatPropertiesNV.html)
  structures in which supported formats and image parameters are
  returned.

## <a href="#_description" class="anchor"></a>Description

If `pImageFormatProperties` is `NULL`, then the number of optical flow
properties supported for the given `physicalDevice` is returned in
`pFormatCount`. Otherwise, `pFormatCount` must point to a variable set
by the user to the number of elements in the `pImageFormatProperties`
array, and on return the variable is overwritten with the number of
values actually written to `pImageFormatProperties`. If the value of
`pFormatCount` is less than the number of optical flow properties
supported, at most `pFormatCount` values will be written to
`pImageFormatProperties`, and `VK_INCOMPLETE` will be returned instead
of `VK_SUCCESS`, to indicate that not all the available values were
returned.

Before creating an image to be used as a optical flow frame, obtain the
supported image creation parameters by querying with
[vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)
and
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
using one of the reported formats and adding
[VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html) to
the `pNext` chain of
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html).

When querying the parameters with
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
for images used for optical flow operations, the
[VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html)::`usage`
field **must** contain one or more of the bits defined in
[VkOpticalFlowUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowUsageFlagBitsNV.html).

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pOpticalFlowImageFormatInfo-parameter"
  id="VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pOpticalFlowImageFormatInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pOpticalFlowImageFormatInfo-parameter  
  `pOpticalFlowImageFormatInfo` **must** be a valid pointer to a valid
  [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pFormatCount-parameter"
  id="VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pFormatCount-parameter"></a>
  VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pFormatCount-parameter  
  `pFormatCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pImageFormatProperties-parameter"
  id="VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pImageFormatProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pImageFormatProperties-parameter  
  If the value referenced by `pFormatCount` is not `0`, and
  `pImageFormatProperties` is not `NULL`, `pImageFormatProperties`
  **must** be a valid pointer to an array of `pFormatCount`
  [VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatPropertiesNV.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_EXTENSION_NOT_PRESENT`

- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_FORMAT_NOT_SUPPORTED`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_FORMAT_B8G8R8A8_UNORM</code>,
<code>VK_FORMAT_R8_UNORM</code> and
<code>VK_FORMAT_G8_B8R8_2PLANE_420_UNORM</code> are initially supported
for images with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#opticalflow-usage"
target="_blank" rel="noopener">optical usage</a>
<code>VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV</code>.</p>
<p><code>VK_FORMAT_R16G16_SFIXED5_NV</code> is initially supported for
images with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#opticalflow-usage"
target="_blank" rel="noopener">optical flow usage</a>
<code>VK_OPTICAL_FLOW_USAGE_OUTPUT_BIT_NV</code>,
<code>VK_OPTICAL_FLOW_USAGE_HINT_BIT_NV</code> and
<code>VK_OPTICAL_FLOW_USAGE_GLOBAL_FLOW_BIT_NV</code>.</p>
<p><code>VK_FORMAT_R8_UINT</code> and <code>VK_FORMAT_R32_UINT</code>
are initially supported for images with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#opticalflow-usage"
target="_blank" rel="noopener">optical flow usage</a>
<code>VK_OPTICAL_FLOW_USAGE_COST_BIT_NV</code>. It is recommended to use
<code>VK_FORMAT_R8_UINT</code> because of the lower bandwidth.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatInfoNV.html),
[VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowImageFormatPropertiesNV.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceOpticalFlowImageFormatsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
