# VK\_KHR\_spirv\_1\_4(3) Manual Page

## Name

VK\_KHR\_spirv\_1\_4 - device extension



## [](#_registered_extension_number)Registered Extension Number

237

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_shader\_float\_controls](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_float_controls.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_spirv_1_4%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_spirv_1_4%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-04-01

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Requires SPIR-V 1.4.

**Contributors**

- Alexander Galazin, Arm
- David Neto, Google
- Jesse Hall, Google
- John Kessenich, Google
- Neil Henning, AMD
- Tom Olson, Arm

## [](#_description)Description

This extension allows the use of SPIR-V 1.4 shader modules. SPIR-V 1.4’s new features primarily make it an easier target for compilers from high-level languages, rather than exposing new hardware functionality.

SPIR-V 1.4 incorporates features that are also available separately as extensions. SPIR-V 1.4 shader modules do not need to enable those extensions with the `OpExtension` opcode, since they are integral parts of SPIR-V 1.4.

SPIR-V 1.4 introduces new floating-point execution mode capabilities, also available via `SPV_KHR_float_controls`. Implementations are not required to support all of these new capabilities; support can be queried using [VkPhysicalDeviceFloatControlsPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFloatControlsPropertiesKHR.html) from the `VK_KHR_shader_float_controls` extension.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SPIRV_1_4_EXTENSION_NAME`
- `VK_KHR_SPIRV_1_4_SPEC_VERSION`

## [](#_issues)Issues

1\. Should we have an extension specific to this SPIR-V version, or add a version-generic query for SPIR-V version? SPIR-V 1.4 does not need any other API changes.

**RESOLVED**: Just expose SPIR-V 1.4.

Most new SPIR-V versions introduce optionally-required capabilities or have implementation-defined limits, and would need more API and specification changes specific to that version to make them available in Vulkan. For example, to support the subgroup capabilities added in SPIR-V 1.3 required introducing [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html) to allow querying the supported group operation categories, maximum supported subgroup size, etc. While we could expose the parts of a new SPIR-V version that do not need accompanying changes generically, we will still end up writing extensions specific to each version for the remaining parts. Thus the generic mechanism will not reduce future spec-writing effort. In addition, making it clear which parts of a future version are supported by the generic mechanism and which cannot be used without specific support would be difficult to get right ahead of time.

2\. Can different stages of the same pipeline use shaders with different SPIR-V versions?

**RESOLVED**: Yes.

Mixing SPIR-V versions 1.0-1.3 in the same pipeline has not been disallowed, so it would be inconsistent to disallow mixing 1.4 with previous versions. SPIR-V 1.4 does not introduce anything that should cause new difficulties here.

3\. Must Vulkan extensions corresponding to SPIR-V extensions that were promoted to core in 1.4 be enabled in order to use that functionality in a SPIR-V 1.4 module?

**RESOLVED**: No, with caveats.

The SPIR-V 1.4 module does not need to declare the SPIR-V extensions, since the functionality is now part of core, so there is no need to enable the Vulkan extension that allows SPIR-V modules to declare the SPIR-V extension. However, when the functionality that is now core in SPIR-V 1.4 is optionally supported, the query for support is provided by a Vulkan extension, and that query can only be used if the extension is enabled.

This applies to any SPIR-V version; specifically for SPIR-V 1.4 this only applies to the functionality from `SPV_KHR_float_controls`, which was made available in Vulkan by `VK_KHR_shader_float_controls`. Even though the extension was promoted in SPIR-V 1.4, the capabilities are still optional in implementations that support `VK_KHR_spirv_1_4`.

A SPIR-V 1.4 module does not need to enable `SPV_KHR_float_controls` in order to use the capabilities, so if the application has *a priori* knowledge that the implementation supports the capabilities, it does not need to enable `VK_KHR_shader_float_controls`. However, if it does not have this knowledge and has to query for support at runtime, it must enable `VK_KHR_shader_float_controls` in order to use [VkPhysicalDeviceFloatControlsPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFloatControlsPropertiesKHR.html).

## [](#_version_history)Version History

- Revision 1, 2019-04-01 (Jesse Hall)
  
  - Internal draft versions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_spirv_1_4)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0