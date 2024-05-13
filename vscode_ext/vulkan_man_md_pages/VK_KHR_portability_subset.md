# VK_KHR_portability_subset(3) Manual Page

## Name

VK_KHR_portability_subset - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

164

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

- **This is a *provisional* extension and **must** be used with caution.
  See the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-provisional-header"
  target="_blank" rel="noopener">description</a> of provisional header
  files for enablement and stability details.**

## <a href="#_contact" class="anchor"></a>Contact

- Bill Hollings <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_portability_subset%5D%20@billhollings%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_portability_subset%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>billhollings</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-07-21

**IP Status**  
No known IP claims.

**Contributors**  
- Bill Hollings, The Brenwill Workshop Ltd.

- Daniel Koch, NVIDIA

- Dzmitry Malyshau, Mozilla

- Chip Davis, CodeWeavers

- Dan Ginsburg, Valve

- Mike Weiblen, LunarG

- Neil Trevett, NVIDIA

- Alexey Knyazev, Independent

## <a href="#_description" class="anchor"></a>Description

The [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
extension allows a non-conformant Vulkan implementation to be built on
top of another non-Vulkan graphics API, and identifies differences
between that implementation and a fully-conformant native Vulkan
implementation.

This extension provides Vulkan implementations with the ability to mark
otherwise-required capabilities as unsupported, or to establish
additional properties and limits that the application should adhere to
in order to guarantee portable behavior and operation across platforms,
including platforms where Vulkan is not natively supported.

The goal of this specification is to document, and make queryable,
capabilities which are required to be supported by a fully-conformant
Vulkan 1.0 implementation, but may be optional for an implementation of
the Vulkan 1.0 Portability Subset.

The intent is that this extension will be advertised only on
implementations of the Vulkan 1.0 Portability Subset, and not on
conformant implementations of Vulkan 1.0. Fully-conformant Vulkan
implementations provide all the required capabilities, and so will not
provide this extension. Therefore, the existence of this extension can
be used to determine that an implementation is likely not fully
conformant with the Vulkan spec.

If this extension is supported by the Vulkan implementation, the
application must enable this extension.

This extension defines several new structures that can be chained to the
existing structures used by certain standard Vulkan calls, in order to
query for non-conformant portable behavior.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDevicePortabilitySubsetPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetPropertiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_PORTABILITY_SUBSET_EXTENSION_NAME`

- `VK_KHR_PORTABILITY_SUBSET_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_PROPERTIES_KHR`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-07-21 (Bill Hollings)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_portability_subset"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
