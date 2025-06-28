# VK\_EXT\_robustness2(3) Manual Page

## Name

VK\_EXT\_robustness2 - device extension



## [](#_registered_extension_number)Registered Extension Number

287

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_robustness2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_robustness2.html) extension

## [](#_contact)Contact

- Liam Middlebrook [\[GitHub\]liam-middlebrook](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_robustness2%5D%20%40liam-middlebrook%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_robustness2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-01-29

**IP Status**

No known IP claims.

**Contributors**

- Liam Middlebrook, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds stricter requirements for how out of bounds reads and writes are handled. Most accesses **must** be tightly bounds-checked, out of bounds writes **must** be discarded, out of bound reads **must** return zero. Rather than allowing multiple possible (0,0,0,x) vectors, the out of bounds values are treated as zero, and then missing components are inserted based on the format as described in [Conversion to RGBA](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-conversion-to-rgba) and [vertex input attribute extraction](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fxvertex-input-extraction).

These additional requirements **may** be expensive on some implementations, and should only be enabled when truly necessary.

This extension also adds support for “null descriptors”, where [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **can** be used instead of a valid handle. Accesses to null descriptors have well-defined behavior, and do not rely on robustness.

## [](#_promotion_to_vk_khr_robustness2)Promotion to `VK_KHR_robustness2`

All functionality in this extension is included in `VK_KHR_robustness2`, with the suffix changed to KHR. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRobustness2FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRobustness2FeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceRobustness2PropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRobustness2PropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_ROBUSTNESS_2_EXTENSION_NAME`
- `VK_EXT_ROBUSTNESS_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_PROPERTIES_EXT`

## [](#_issues)Issues

1. Why do [VkPhysicalDeviceRobustness2PropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRobustness2PropertiesEXT.html)::`robustUniformBufferAccessSizeAlignment` and [VkPhysicalDeviceRobustness2PropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRobustness2PropertiesEXT.html)::`robustStorageBufferAccessSizeAlignment` exist?

**RESOLVED**: Some implementations cannot efficiently tightly bounds-check all buffer accesses. Rather, the size of the bound range is padded to some power of two multiple, up to 256 bytes for uniform buffers and up to 4 bytes for storage buffers, and that padded size is bounds-checked. This is sufficient to implement D3D-like behavior, because D3D only allows binding whole uniform buffers or ranges that are a multiple of 256 bytes, and D3D raw and structured buffers only support 32-bit accesses.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2019-11-01 (Jeff Bolz, Liam Middlebrook)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_robustness2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0