# VK\_KHR\_shader\_integer\_dot\_product(3) Manual Page

## Name

VK\_KHR\_shader\_integer\_dot\_product - device extension



## [](#_registered_extension_number)Registered Extension Number

281

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_integer\_dot\_product](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_integer_dot_product.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_integer_dot_product%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_integer_dot_product%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_integer\_dot\_product](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_integer_dot_product.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-16

**Interactions and External Dependencies**

- This extension interacts with `VK_KHR_shader_float16_int8`.

**IP Status**

No known IP claims.

**Contributors**

- Kévin Petit, Arm Ltd.
- Jeff Bolz, NVidia
- Spencer Fricke, Samsung
- Jesse Hall, Google
- John Kessenich, Google
- Graeme Leese, Broadcom
- Einar Hov, Arm Ltd.
- Stuart Brady, Arm Ltd.
- Pablo Cascon, Arm Ltd.
- Tobias Hector, AMD
- Jeff Leger, Qualcomm
- Ruihao Zhang, Qualcomm
- Pierre Boudier, NVidia
- Jon Leech, The Khronos Group
- Tom Olson, Arm Ltd.

## [](#_description)Description

This extension adds support for the integer dot product SPIR-V instructions defined in SPV\_KHR\_integer\_dot\_product. These instructions are particularly useful for neural network inference and training but find uses in other general-purpose compute applications as well.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderIntegerDotProductFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderIntegerDotProductFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderIntegerDotProductPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderIntegerDotProductPropertiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_INTEGER_DOT_PRODUCT_EXTENSION_NAME`
- `VK_KHR_SHADER_INTEGER_DOT_PRODUCT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_PROPERTIES_KHR`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the KHR suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`DotProductInputAllKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DotProductInputAll)
- [`DotProductInput4x8BitKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DotProductInput4x8Bit)
- [`DotProductInput4x8BitPackedKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DotProductInput4x8BitPacked)
- [`DotProductKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DotProduct)

## [](#_version_history)Version History

- Revision 1, 2021-06-16 (Kévin Petit)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_integer_dot_product).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0