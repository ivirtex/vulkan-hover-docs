# VK_KHR_uniform_buffer_standard_layout(3) Manual Page

## Name

VK_KHR_uniform_buffer_standard_layout - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

254

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Graeme Leese <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_uniform_buffer_standard_layout%5D%20@gnl21%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_uniform_buffer_standard_layout%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gnl21</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-01-25

**Contributors**  
- Graeme Leese, Broadcom

- Jeff Bolz, NVIDIA

- Tobias Hector, AMD

- Faith Ekstrand, Intel

- Neil Henning, AMD

## <a href="#_description" class="anchor"></a>Description

This extension enables tighter array and struct packing to be used with
uniform buffers.

It modifies the alignment rules for uniform buffers, allowing for
tighter packing of arrays and structures. This allows, for example, the
std430 layout, as defined in
[GLSL](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.pdf)
to be supported in uniform buffers.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_EXTENSION_NAME`

- `VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-01-25 (Graeme Leese)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_uniform_buffer_standard_layout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
