# vkEnumerateDeviceExtensionProperties(3) Manual Page

## Name

vkEnumerateDeviceExtensionProperties - Returns properties of available physical device extensions



## [](#_c_specification)C Specification

To query the extensions available to a given physical device, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkEnumerateDeviceExtensionProperties(
    VkPhysicalDevice                            physicalDevice,
    const char*                                 pLayerName,
    uint32_t*                                   pPropertyCount,
    VkExtensionProperties*                      pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be queried.
- `pLayerName` is either `NULL` or a pointer to a null-terminated UTF-8 string naming the layer to retrieve extensions from.
- `pPropertyCount` is a pointer to an integer related to the number of extension properties available or queried, and is treated in the same fashion as the [vkEnumerateInstanceExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumerateInstanceExtensionProperties.html)::`pPropertyCount` parameter.
- `pProperties` is either `NULL` or a pointer to an array of [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html) structures.

## [](#_description)Description

When `pLayerName` parameter is `NULL`, only extensions provided by the Vulkan implementation or by implicitly enabled layers are returned. When `pLayerName` is the name of a layer, the device extensions provided by that layer are returned.

Implementations **must** not advertise any pair of extensions that cannot be enabled together due to behavioral differences, or any extension that cannot be enabled against the advertised version.

If the `VK_KHR_ray_tracing_pipeline` extension is advertised as supported by this query, the `VK_KHR_pipeline_library` extension **must** also be supported.

Implementations claiming support for the [Roadmap 2022](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#roadmap-2022) profile **must** advertise the `VK_KHR_global_priority` extension in `pProperties`.

Implementations claiming support for the [Roadmap 2024](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#roadmap-2024) profile **must** advertise the following extensions in `pProperties`:

- [VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html)
- [VK\_KHR\_load\_store\_op\_none](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_load_store_op_none.html)
- [VK\_KHR\_shader\_quad\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_quad_control.html)
- [VK\_KHR\_shader\_maximal\_reconvergence](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_maximal_reconvergence.html)
- [VK\_KHR\_shader\_subgroup\_uniform\_control\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_uniform_control_flow.html)
- [VK\_KHR\_shader\_subgroup\_rotate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_rotate.html)
- [VK\_KHR\_shader\_float\_controls2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_float_controls2.html)
- [VK\_KHR\_shader\_expect\_assume](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_expect_assume.html)
- [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html)
- [VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html)
- [VK\_KHR\_index\_type\_uint8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_index_type_uint8.html)
- [VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html)
- [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)
- [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html)

Note

Due to platform details on Android, `vkEnumerateDeviceExtensionProperties` may be called with `physicalDevice` equal to `NULL` during layer discovery. This behavior will only be observed by layer implementations, and not the underlying Vulkan driver.

Valid Usage (Implicit)

- [](#VUID-vkEnumerateDeviceExtensionProperties-physicalDevice-parameter)VUID-vkEnumerateDeviceExtensionProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkEnumerateDeviceExtensionProperties-pLayerName-parameter)VUID-vkEnumerateDeviceExtensionProperties-pLayerName-parameter  
  If `pLayerName` is not `NULL`, `pLayerName` **must** be a null-terminated UTF-8 string
- [](#VUID-vkEnumerateDeviceExtensionProperties-pPropertyCount-parameter)VUID-vkEnumerateDeviceExtensionProperties-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkEnumerateDeviceExtensionProperties-pProperties-parameter)VUID-vkEnumerateDeviceExtensionProperties-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertyCount` [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html) structures

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_LAYER_NOT_PRESENT`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtensionProperties.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkEnumerateDeviceExtensionProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0