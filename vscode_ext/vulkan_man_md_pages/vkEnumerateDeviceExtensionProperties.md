# vkEnumerateDeviceExtensionProperties(3) Manual Page

## Name

vkEnumerateDeviceExtensionProperties - Returns properties of available
physical device extensions



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the extensions available to a given physical device, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkEnumerateDeviceExtensionProperties(
    VkPhysicalDevice                            physicalDevice,
    const char*                                 pLayerName,
    uint32_t*                                   pPropertyCount,
    VkExtensionProperties*                      pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be queried.

- `pLayerName` is either `NULL` or a pointer to a null-terminated UTF-8
  string naming the layer to retrieve extensions from.

- `pPropertyCount` is a pointer to an integer related to the number of
  extension properties available or queried, and is treated in the same
  fashion as the
  [vkEnumerateInstanceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html)::`pPropertyCount`
  parameter.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structures.

## <a href="#_description" class="anchor"></a>Description

When `pLayerName` parameter is `NULL`, only extensions provided by the
Vulkan implementation or by implicitly enabled layers are returned. When
`pLayerName` is the name of a layer, the device extensions provided by
that layer are returned.

Implementations **must** not advertise any pair of extensions that
cannot be enabled together due to behavioral differences, or any
extension that cannot be enabled against the advertised version.

Implementations claiming support for the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#roadmap-2022"
target="_blank" rel="noopener">Roadmap 2022</a> profile **must**
advertise the [`VK_KHR_global_priority`](VK_KHR_global_priority.html)
extension in `pProperties`.

Implementations claiming support for the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#roadmap-2024"
target="_blank" rel="noopener">Roadmap 2024</a> profile **must**
advertise the following extensions in `pProperties`:

- [VK_KHR_dynamic_rendering_local_read](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering_local_read.html)

- [VK_KHR_load_store_op_none](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_load_store_op_none.html)

- [VK_KHR_shader_quad_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_quad_control.html)

- [VK_KHR_shader_maximal_reconvergence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_maximal_reconvergence.html)

- [VK_KHR_shader_subgroup_uniform_control_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_subgroup_uniform_control_flow.html)

- [VK_KHR_shader_subgroup_rotate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_subgroup_rotate.html)

- [VK_KHR_shader_float_controls2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float_controls2.html)

- [VK_KHR_shader_expect_assume](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_expect_assume.html)

- [VK_KHR_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_line_rasterization.html)

- [VK_KHR_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vertex_attribute_divisor.html)

- [VK_KHR_index_type_uint8](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_index_type_uint8.html)

- [VK_KHR_map_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_map_memory2.html)

- [VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html)

- [VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html)

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Due to platform details on Android,
<code>vkEnumerateDeviceExtensionProperties</code> may be called with
<code>physicalDevice</code> equal to <code>NULL</code> during layer
discovery. This behavior will only be observed by layer implementations,
and not the underlying Vulkan driver.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a
  href="#VUID-vkEnumerateDeviceExtensionProperties-physicalDevice-parameter"
  id="VUID-vkEnumerateDeviceExtensionProperties-physicalDevice-parameter"></a>
  VUID-vkEnumerateDeviceExtensionProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkEnumerateDeviceExtensionProperties-pLayerName-parameter"
  id="VUID-vkEnumerateDeviceExtensionProperties-pLayerName-parameter"></a>
  VUID-vkEnumerateDeviceExtensionProperties-pLayerName-parameter  
  If `pLayerName` is not `NULL`, `pLayerName` **must** be a
  null-terminated UTF-8 string

- <a
  href="#VUID-vkEnumerateDeviceExtensionProperties-pPropertyCount-parameter"
  id="VUID-vkEnumerateDeviceExtensionProperties-pPropertyCount-parameter"></a>
  VUID-vkEnumerateDeviceExtensionProperties-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkEnumerateDeviceExtensionProperties-pProperties-parameter"
  id="VUID-vkEnumerateDeviceExtensionProperties-pProperties-parameter"></a>
  VUID-vkEnumerateDeviceExtensionProperties-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html) structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_LAYER_NOT_PRESENT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEnumerateDeviceExtensionProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
