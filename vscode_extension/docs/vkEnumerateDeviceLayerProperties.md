# vkEnumerateDeviceLayerProperties(3) Manual Page

## Name

vkEnumerateDeviceLayerProperties - Returns properties of available
physical device layers



## <a href="#_c_specification" class="anchor"></a>C Specification

To enumerate device layers, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkEnumerateDeviceLayerProperties(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkLayerProperties*                          pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be queried.

- `pPropertyCount` is a pointer to an integer related to the number of
  layer properties available or queried.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html) structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of layer properties
available is returned in `pPropertyCount`. Otherwise, `pPropertyCount`
**must** point to a variable set by the application to the number of
elements in the `pProperties` array, and on return the variable is
overwritten with the number of structures actually written to
`pProperties`. If `pPropertyCount` is less than the number of layer
properties available, at most `pPropertyCount` structures will be
written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`,
to indicate that not all the available properties were returned.

The list of layers enumerated by `vkEnumerateDeviceLayerProperties`
**must** be exactly the sequence of layers enabled for the instance. The
members of `VkLayerProperties` for each enumerated layer **must** be the
same as the properties when the layer was enumerated by
`vkEnumerateInstanceLayerProperties`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Due to platform details on Android,
<code>vkEnumerateDeviceLayerProperties</code> may be called with
<code>physicalDevice</code> equal to <code>NULL</code> during layer
discovery. This behavior will only be observed by layer implementations,
and not the underlying Vulkan driver.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a
  href="#VUID-vkEnumerateDeviceLayerProperties-physicalDevice-parameter"
  id="VUID-vkEnumerateDeviceLayerProperties-physicalDevice-parameter"></a>
  VUID-vkEnumerateDeviceLayerProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkEnumerateDeviceLayerProperties-pPropertyCount-parameter"
  id="VUID-vkEnumerateDeviceLayerProperties-pPropertyCount-parameter"></a>
  VUID-vkEnumerateDeviceLayerProperties-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a href="#VUID-vkEnumerateDeviceLayerProperties-pProperties-parameter"
  id="VUID-vkEnumerateDeviceLayerProperties-pProperties-parameter"></a>
  VUID-vkEnumerateDeviceLayerProperties-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html) structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEnumerateDeviceLayerProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
