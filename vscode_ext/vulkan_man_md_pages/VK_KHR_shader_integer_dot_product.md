# VK_KHR_shader_integer_dot_product(3) Manual Page

## Name

VK_KHR_shader_integer_dot_product - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

281

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_integer_dot_product](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_integer_dot_product.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Kevin Petit <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_integer_dot_product%5D%20@kpet%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_integer_dot_product%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kpet</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_shader_integer_dot_product](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_integer_dot_product.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-06-16

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_KHR_shader_float16_int8`](VK_KHR_shader_float16_int8.html).

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

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the integer dot product SPIR-V
instructions defined in SPV_KHR_integer_dot_product. These instructions
are particularly useful for neural network inference and training but
find uses in other general-purpose compute applications as well.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderIntegerDotProductFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderIntegerDotProductFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderIntegerDotProductPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderIntegerDotProductPropertiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_INTEGER_DOT_PRODUCT_EXTENSION_NAME`

- `VK_KHR_SHADER_INTEGER_DOT_PRODUCT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_PROPERTIES_KHR`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
KHR suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DotProductInputAllKHR"
  target="_blank" rel="noopener"><code>DotProductInputAllKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DotProductInput4x8BitKHR"
  target="_blank" rel="noopener"><code>DotProductInput4x8BitKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DotProductInput4x8BitPackedKHR"
  target="_blank"
  rel="noopener"><code>DotProductInput4x8BitPackedKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DotProductKHR"
  target="_blank" rel="noopener"><code>DotProductKHR</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-06-16 (Kévin Petit)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_integer_dot_product"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
